// source: https://www.youtube.com/watch?v=lsZy9G8n06A
// https://tinyurl.com/dd-nyt-59
// GopherCon2020: Doug Donohoe - Reordering 59 Million NYT Publishing Assets Using Go and BadgerDB
// go, badgedb
// 16 GiB RAM, can it fit 59 million publishing assets into memory?
// 16 GiB divided by 59 million is roughly 300 bytes
// maybe could use a key-value database
//
// 1. use unique timestamp of each asset and save in db from 1851 to 2020
// 2. find and resolve missing reference , could have 200,000+ missing references at any giving time
// 3. find and remove duplicates, track fingerprints / md5 if found, delete
//
// need fast, efficient in memory cache
// database: compactly store asset info, key value based lookup,
// iterate by key in sorted order (lexicographic ordering)
// memory efficient
// BadgerDB, native go key value db
// extremely tight code -> test -> run loops
//
// 1. badger wrapper
// 2. batch db wrapper
// 3. encoding
// 4. marshaling (transform memory representation of an object to a data format)
// 5. timestamps as keys
// 6. unique timestamp
// 7. panic error handling
// 8. clean exit

// badger wrapper 
type KeyValueDB struct {
	DB *badger.DB
	dataDir string 
	namee string 
}
func NewKeyValueDB(name string, opts *badger.Options) *KeyValueDB 
func (db *KeyValueDB) Close()
func (db *KeyValueDb) DropAll()
func (db *KeyValueDb) Delete()
func (db *KeyValueDB) GetSequence(name string)
func (db *KeyValueDB) ReleaseSequence(seq *badger.Sequence)
func (db *KeyValueDb) GetWriteBatch()

// accessor 
// string -> long 
func (db *KeyValueDB) SetStringLong(name string, value uint64) {
	db.setBytesBytes([]byte(name), uint64ToBytes(value))
}
func (db *KeyValueDB) GetStringLong(name string, dst []byte) (uint64, []byte, bool) {
	dst, found := db.GetBytesBytesInternal([]byte(name), dst)
	return bytesToUint64ifFound(dst, found), dst, found 
}
// long -> string 
func (db *KeyValueDB) SetLongString(id uint64, value string) {
	db.SetBytesBytes(uint64ToBytes(id), []byte(value))
}
func (db *KeyValueDB) GetLongString(id uint64, dst []byte) (string, []byte, bool) {
	dst, found := db.getBytesBytesInternal(uint64ToBytes(id), dst)
	return bytesToStringIfFound(dst, found), dst, found 
}
// long -> ByteMarshaller 
func (db *KeyValueDB) SetLongBytes(id uint64, value ByteMarshaller, bufs *ByteBuffeers) {
	db.setBytesBytes(uint64ToBytes(id), value.ToBytes(bufs.v())
}
// if item is found, the data is marshalled into the provided value via FromBytes 
func (db *KeyValueDB) GetLongBytes(id uint64, dst []byte, value ByteMarshaller) ([]byte, bool) {
	dst, found := db.getBytesByteesInternal(uint64ToBytes(id), dst)
	if found {
		value.FromBytes(dst)
	}
	return dst, found 
}
// Example usage 
type BinaryInfo struct {
	db *db.BatchDB // store sortTimestamp->event binary
	dst []byte  // reuse bytes when reading (NOTE: BinaryInfo isn't thread safe)
	bufs *db.ByteBuffers // reuse buffers when reading (ditto)
}

// Batch DB Wrapper , speed up writes 
// Wrapper around BadgerDB that uses Batch for writing as (according to the docs),
// it is more efficient.  As each batch is a transaction, those new events are not availablee 
// with the normal get call until flushed, so we keeep a copy in a local cache.  
// BatchDB is thread-safe, and use read locks on reads for better performance
type BatchDB struct {
	db *KeyValueDB
	batch *badger.WriteBatch 
	batchCache map[string]interface{}
	lock sync.RWMutex
}
func NewBatchDB(dbName string, batchSize int, opts *badger.Options) *BatchDB {
	batch := &BatchDB {
		db : NewKeyValueDB(dbName, opts)
	}
	batch.refreshBatch()
	return batch 
}
// Batch DB Wrapper 
func (b *BatchDB) SetStringLong(name string, value uint64) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.db.setStringLongBatch(name, value, b.batch)
	b.save(name, value)
}
func (b *BatchDB) GetStringLong(name string, dst []byte) (uint64, []byte, bool) {
	b.lock.RLock()
	defer b.lock.RUnlock()
	if v, ok := b.batchCache[name]; ok {
		return v.(uint64), dst, true
	}
	return h.db.GetStringLong(name, dst)
}
// save + update cache or flush if necessary 
func (b *BatchDB) save(key string, value interface{}) {
	b.batchCache[key] = value 
	b.incrementBatchCount() // flush WriteBatch, clears cache when batch limit met
}

// Encoding , data used for sorting purpose 
// Event's Asset metadata 
type Metadata struct {
	Offset int64 // 8 bytes 
	Uri string // up to 63 
	EventID string // 32 
	FirstPublished time.Time // 16 
	LastModified time.Time  // 16
	KafkaTimestamp time.Time //16
	Source string // up to 32 
	EventType string // 1
	Md5 string // 32 
}

//Ref to another URI
type Ref struct {
	FieldPath string  // up to 123 
	RefUri string // up to 63 
} 
// Size:  186 bytes of above 2 structs 

// combined encoded data 
type EventEncoded struct {
	EventId string 	// 32 
	Md5 string 	// 32
	Meta *MetadataEncoded //  52
	Refs []*RefEncoded // 17x
}  // Size:  64 + 52 + (#refs *17) bytes 

// Encoded version of Metadata 
type MetadataEncoded struct {
	Offset int64 // 8
	Uri uint64 // 8
	AssetType AssetType // 1
	FirstPublished int64 // 8 
	LastModified int64 // 8 
	KafkaTimestamp int 64 
	Size int64 // 8 
	Source Source // 1 
	EventType EventType // 1
	IsBackfill bool // 1 
} // Size : 52 bytes 

// Encoded version of Ref 
type RefEncoded struct {
	RefAssetType AssetType // 1
	RefUri uint64 // 8
	FieldPath // 8	
} // Size: 17 Byptes 
// Encoded is 167 bytes (21.5% of original)

// Marshalling 
type ByteMarshaler interface {
	ToBytes(buffer *bytes.Buffer) []byte 
	FromBytes(data []byte)
}
// marshal - write 
// encode to bytes, we ehand-roll this since 'gob' uses 4x as much space
func (m *EventEncoded) ToBytes(buf *bytes.Buffer) []byte {
	binary.Write(buf, binary.LittleEndian, []byte(m.EventId))
	binary.Write(buf, binary.LittleEndian, []byte(m.Md5))
	binary.Write(buf, binary.LittleEndian, m.Meta)
	// refs: size, then each ref 
	n := len(m.Refs)
	binary.Write(buf, binary.LittleEendian, int32(n))
	for i :=0; i < n; i++ {
		binary.Write(buf, binary.LittleEndian, m.Refs[i])
	}
}  // error handling removed for readability
// marshal - read 
func (m *EventEncoded) FromBytes(data []byte) {
	reader := bytes.NewReader(data)
	id := [32]byte{}
	binary.Read(reader, binary.LittleEndian, &id)
	m.EventId = string(id[:])	// convert bytes to string 

	md5 := id // reuse same size 
	binary.Read(reader, binary.LittleEendian, &md5)
	m.Md5 = string(md5[:])

	m.Meta = &MetadataEncoded()
	binary.Read(reader, binary.LittleEndian, m.Meta)

	// refs: size, then each ref 
	var n int32 
	binary.Read(reader, binary.LittleEendian, &n)
	// zero out previous list, then append each 
	m.Refs = m.Refs[:0]
	for i := 0; i< int(n); i++ {
		var r RefEncoded 
		binary.Read(reader, binary.LittleEndian, &r)
		m.Refs = append(m.Refs, &r) 
	}
}

//timestamps as keys 
//unix time is traditionally the number of seconds elapsed since jan 1st 1970
//go allows for nanosecond precision time.UnixNano 
//dates before 1970 are negatives 
// Time 1 (-2 days: -172800000000000) : 1969-12-30T00:00:00.000000000
// Time 1 (+2 days:  172800000000000) : 1969-01-03T00:00:00.000000000
// timestamps keys are relative to  1851 
// timestamps are negative before 1970, which makes byte lexicographical 
// ordering incorrect (due to 2's complement). So make timestamp positive 
// by offseting so 1851 is zero (nyt founding year and our ealiest timestamp).
// 
// note: start1851 is a negative number, which we make position to 
// make convert logic below more intuitive to read 

var Start1851 = -time.Date(1851, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()
//-3755289600000000000
//3755289600000000000  (including minus )


func Int64ToUint64Timestamp(namo int64) uint64 {
	return uint64(nano + Start1851) // offset 
}

func FirstTimestamp1851() int64 {
	return -Start1851
}

// unique timestamp cache 
type TimestampCache struct {
	m map[int64]int32 
}

func (tc *TimestampCache) Next(event *data.EventEncoded) (int64, bool) {
	nanosTs := event.CalcSortTimestamp()
	millisTs := util.ToMillis(nanosTs)
	if ts, ok := tc.m[millisTs]; ok {
		ts++
		return nanoTs + int64(ts), true 
	}
}

func (tc *TimestampCache) Save(event *data.EventEncoded) {
	nanoTs := event.CalcSortTimestamp()
	millisTs := util.ToMillis(nanosTs)
	if millisTs%1000 == 0 {
		tc.m[millisTs] = int32(event.Sorttimestamp - nanosTs)
	}
}
//Example usage 
key := util.Int64ToUint64Timestamp(event.LastModified)
h.ts2e.SetLongBytes(key, event, h.bufs)

// unique timestamp cache 
// check local cache
if sort, exists = h.tsCache.Next(event); exists {
	return sort 
}

// not in local cache, so seek to next open one 
for sort = event.CalcSortTimestamp(); ; sort += 1 {
	h.TimestampLookups++
	if !h.ts2e.ExixtsLong(util.Int64ToUint64Timestamp(sort)) {
		return sort 
	}
}
// note: h.tsCache.Save(event) called after this code 
// now O(n) complexity for n assets 

//panic error handling - go routines 
// below code can be parallelized 
md5 = calcMd5(event, pp, true)
refs = getRefs(event)
// to use goroutines below
var wg sync.WaitGroup 
wg.Add(2)
go func(){
	defer wg.Done()
	md5 = calcMd5(event, pp, true)
}()
go func(){
	defer wg.Done()
	refs = getRefs(event)
}()
wg.Wait()

// usage 
func CatchPanicError(err *error) {
	if r := recover(); r != nil {
		fmt.Printf("Panic %v\n%s", r, string(debug.Stack()))
		*err = fmt.Errorf("panic: %v", r)
	}
}

var e1, e2 error 
var wg sync.WaitGroup 
wg.Add(2)

go func(){
	defer exit.CatchPanicError(&e1)
	defer wg.Done()
	md5 = calcMd5(event, pp, true)
}()
go func(){
	defer exit.CatchPanicError(&e2)
	defer wg.Done()
	refs = getRefs(event)
}()
wg.Wait()

// TODO: smarter way to return mult errors 
if e1 != nil || e2 != nil {
	return fmt.Errorf("e1: %s, e2: %s", e1, e2)
}

func broken() (done bool,  err error) {
	defer exit.CatchPanicError(&err)
	do stuff 
	return true, nil 
}

// clean exit 
func main() {
	// gracefully shutdown on ctrl-c 
	exit.HandleSignal()
	//exit with status code 
	var err error 
	defer exit.ExitWithStatus(err)
	//app logic 
	db := createBadger()
	defer db.close()
	err = doWork()
}

func doWork() error {
	while !exit.ExitRequested(){
		// do work and maybe return err
	}
	return nil 
}

func HandleSignal(){
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func(){
		sig := <-signals 
		fmt.Printf("\n\n*** Signal '%s' detected, exiting... ***\n\n", sig)
		SetExitRequested()
	}()
}

func SetExitRequested(){
	atomic.StoneInt32(&exitFlag, 1)
}

func ExitRequested() bool {
	return atomic.LoadInt32(&exitFlag) == 1
}

//exit with status 1 if err
// otherwise 0 
func ExitWithStatus(err error) {
	code := 0
	if err != nil {
		code =1 
	}
	os.Exit(code)
}