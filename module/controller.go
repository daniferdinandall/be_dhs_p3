package namapackage

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aiteung/atdb"
	model "github.com/daniferdinandall/be_dhs_p3/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

// Presensi
var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "db_dhs_tb",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

// DHS

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

// dhs
func InsertDHS(db *mongo.Database, mahasiswa model.Mahasiswa, mata_kuliah []model.NilaiMataKuliah) (insertedID primitive.ObjectID, err error) {
	dhs := bson.M{
		"mahasiswa":   mahasiswa,
		"mata_kuliah": mata_kuliah,
		"created_at":  primitive.NewDateTimeFromTime(time.Now().UTC()),
	}
	// dhs.Mahasiswa = mahasiswa
	// dhs.MataKuliah = mata_kuliah
	// dhs.CreatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	// return InsertOneDoc("db_dhs_tb", "dhs", dhs)
	result, err := db.Collection("dhs").InsertOne(context.Background(), dhs)
	if err != nil {
		fmt.Printf("InsertPresensi: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetDhsFromNPM(db *mongo.Database, npm int) (dhs model.Dhs) {

	data_dhs := MongoConnect("db_dhs_tb").Collection("dhs")
	filter := bson.M{"mahasiswa.npm": npm}
	err := data_dhs.FindOne(context.TODO(), filter).Decode(&dhs)
	if err != nil {
		fmt.Printf("GetDhsFromNPM: %v\n", err)
	}
	return dhs
}

func GetDhsFromID(db *mongo.Database, _id primitive.ObjectID) (hasil model.Dhs, errs error) {
	data := db.Collection("dhs")
	filter := bson.M{"_id": _id}
	err := data.FindOne(context.TODO(), filter).Decode(&hasil)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return hasil, fmt.Errorf("no data found for ID %s", _id)
		}
		return hasil, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return hasil, nil
}

func GetDhsAll(db *mongo.Database) (dhs []model.Dhs) {
	data_dhs := db.Collection("dhs")
	filter := bson.D{}
	// var results []Dhs
	cur, err := data_dhs.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetDhsFromNPM: %v\n", err)
	}
	err = cur.All(context.TODO(), &dhs)
	if err != nil {
		fmt.Println(err)
	}
	return dhs
}

// mahasiswa
func InsertMhs(npm int, nama string, fakultas model.Fakultas, dosen model.Dosen, programStudi model.ProgramStudi) (InsertedID interface{}) {
	var mhs model.Mahasiswa
	mhs.Npm = npm
	mhs.Nama = nama
	mhs.Fakultas = fakultas
	mhs.DosenWali = dosen
	mhs.ProgramStudi = programStudi
	// mhs.CreatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	return InsertOneDoc("db_dhs_tb", "mahasiswa", mhs)
}

func GetMhsFromNPM(npm int) (mhs model.Mahasiswa) {
	data_dhs := MongoConnect("db_dhs_tb").Collection("mahasiswa")
	filter := bson.M{"npm": npm}
	err := data_dhs.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		fmt.Printf("GetMhsFromNPM: %v\n", err)
	}
	return mhs
}

func GetMhsAll() (mhs []model.Mahasiswa) {
	data_mhs := MongoConnect("db_dhs_tb").Collection("mahasiswa")
	filter := bson.D{}
	// var results []mhs
	cur, err := data_mhs.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetmhsFromNPM: %v\n", err)
	}
	err = cur.All(context.TODO(), &mhs)
	if err != nil {
		fmt.Println(err)
	}
	return mhs
}

// dosen
func InsertDosen(kode string, nama string, hp string) (InsertedID interface{}) {
	var dosen model.Dosen
	dosen.KodeDosen = kode
	dosen.Nama = nama
	dosen.PhoneNumber = hp
	return InsertOneDoc("db_dhs_tb", "dosen", dosen)
}

func GetDosenFromKodeDosen(kode string) (dosen model.Dosen) {
	data_dhs := MongoConnect("db_dhs_tb").Collection("dosen")
	filter := bson.M{"kode_dosen": kode}
	err := data_dhs.FindOne(context.TODO(), filter).Decode(&dosen)
	if err != nil {
		fmt.Printf("GetDosenFromKodeDosen: %v\n", err)
	}
	return dosen
}

func GetDosenAll() (dosen []model.Dosen) {
	data_mhs := MongoConnect("db_dhs_tb").Collection("dosen")
	filter := bson.D{}
	// var results []mhs
	cur, err := data_mhs.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetAllDosen: %v\n", err)
	}
	err = cur.All(context.TODO(), &dosen)
	if err != nil {
		fmt.Println(err)
	}
	return dosen
}

// matkul
func InsertMatkul(kode string, nama string, sks int, dosen model.Dosen) (InsertedID interface{}) {
	var matkul model.MataKuliah
	matkul.KodeMatkul = kode
	matkul.Nama = nama
	matkul.Sks = sks
	matkul.Dosen = dosen
	return InsertOneDoc("db_dhs_tb", "matkul", matkul)
}

func GetMatkulFromKodeMatkul(kode string) (matkul model.MataKuliah) {
	data_dhs := MongoConnect("db_dhs_tb").Collection("matkul")
	filter := bson.M{"kode_matkul": kode}
	err := data_dhs.FindOne(context.TODO(), filter).Decode(&matkul)
	if err != nil {
		fmt.Printf("GetMatkulFromKodeMatkul: %v\n", err)
	}
	return matkul
}

func GetMatkulAll() (matkul []model.MataKuliah) {
	data_mhs := MongoConnect("db_dhs_tb").Collection("matkul")
	filter := bson.D{}
	// var results []mhs
	cur, err := data_mhs.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetMatkulAll: %v\n", err)
	}
	err = cur.All(context.TODO(), &matkul)
	if err != nil {
		fmt.Println(err)
	}
	return matkul
}

// fakultas
func InsertFakultas(kode string, nama string) (InsertedID interface{}) {
	var fakultas model.Fakultas
	fakultas.KodeFakultas = kode
	fakultas.Nama = nama
	return InsertOneDoc("db_dhs_tb", "fakultas", fakultas)
}

func GetFakultasFromKodeFakultas(kode string) (fakultas model.MataKuliah) {
	data_dhs := MongoConnect("db_dhs_tb").Collection("fakultas")
	filter := bson.M{"kode_fakultas": kode}
	err := data_dhs.FindOne(context.TODO(), filter).Decode(&fakultas)
	if err != nil {
		fmt.Printf("GetFakultasFromKodeFakultas: %v\n", err)
	}
	return fakultas
}

func GetFakultasAll() (fakultas []model.MataKuliah) {
	data_mhs := MongoConnect("db_dhs_tb").Collection("fakultas")
	filter := bson.D{}
	// var results []mhs
	cur, err := data_mhs.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetFakultasAll: %v\n", err)
	}
	err = cur.All(context.TODO(), &fakultas)
	if err != nil {
		fmt.Println(err)
	}
	return fakultas
}

// prodi
func InsertProdi(kode string, nama string) (InsertedID interface{}) {
	var programStudi model.ProgramStudi
	programStudi.KodeProgramStudi = kode
	programStudi.Nama = nama
	return InsertOneDoc("db_dhs_tb", "programStudi", programStudi)
}

func GetProdiFromKodeProdi(kode string) (programStudi model.MataKuliah) {
	data_dhs := MongoConnect("db_dhs_tb").Collection("programStudi")
	filter := bson.M{"kode_programStudi": kode}
	err := data_dhs.FindOne(context.TODO(), filter).Decode(&programStudi)
	if err != nil {
		fmt.Printf("GetProdiFromKodeProdi: %v\n", err)
	}
	return programStudi
}

func GetProdiAll() (programStudi []model.MataKuliah) {
	data_mhs := MongoConnect("db_dhs_tb").Collection("programStudi")
	filter := bson.D{}
	// var results []mhs
	cur, err := data_mhs.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetProdiAll: %v\n", err)
	}
	err = cur.All(context.TODO(), &programStudi)
	if err != nil {
		fmt.Println(err)
	}
	return programStudi
}
