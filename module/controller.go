package namapackage

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
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

func GetDhsFromNPM(db *mongo.Database, npm int) (dhs model.Dhs, errs error) {

	data := db.Collection("dhs")
	filter := bson.M{"mahasiswa.npm": npm}
	err := data.FindOne(context.TODO(), filter).Decode(&dhs)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return dhs, fmt.Errorf("no data found for npm %s", strconv.Itoa(npm))
		}
		return dhs, fmt.Errorf("error retrieving data for npm %s: %s", strconv.Itoa(npm), err.Error())
	}
	return dhs, nil
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

func UpdateDhsById(db *mongo.Database, id primitive.ObjectID, mahasiswa model.Mahasiswa, mata_kuliah []model.NilaiMataKuliah) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"mahasiswa":   mahasiswa,
			"mata_kuliah": mata_kuliah,
			"created_at":  primitive.NewDateTimeFromTime(time.Now().UTC()),
		},
	}
	result, err := db.Collection("dhs").UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdatePresensi: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func DeleteDhsByID(db *mongo.Database, _id primitive.ObjectID) error {
	karyawan := db.Collection("dhs")
	filter := bson.M{"_id": _id}

	result, err := karyawan.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// mahasiswa
func InsertMhs(db *mongo.Database, npm int, nama string, fakultas model.Fakultas, dosen model.Dosen, programStudi model.ProgramStudi) (insertedID primitive.ObjectID, err error) {
	var mhs model.Mahasiswa
	mhs.Npm = npm
	mhs.Nama = nama
	mhs.Fakultas = fakultas
	mhs.DosenWali = dosen
	mhs.ProgramStudi = programStudi
	// mhs.CreatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	// return InsertOneDoc("db_dhs_tb", "mahasiswa", mhs)
	result, err := db.Collection("mahasiswa").InsertOne(context.Background(), mhs)
	if err != nil {
		fmt.Printf("InsertPresensi: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetMhsFromNPM(db *mongo.Database, npm int) (mhs model.Mahasiswa, errs error) {
	data_dhs := db.Collection("mahasiswa")
	filter := bson.M{"npm": npm}
	err := data_dhs.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return mhs, fmt.Errorf("no data found for npm %s", strconv.Itoa(npm))
		}
		return mhs, fmt.Errorf("error retrieving data for npm %s: %s", strconv.Itoa(npm), err.Error())
	}
	return mhs, nil
}

func GetMhsFromID(db *mongo.Database, _id primitive.ObjectID) (hasil model.Mahasiswa, errs error) {
	data := db.Collection("mahasiswa")
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

func GetMhsAll(db *mongo.Database) (mhs []model.Mahasiswa) {
	data_mhs := db.Collection("mahasiswa")
	filter := bson.D{}
	// var results []mhs
	cur, err := data_mhs.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetMhsFromNPM: %v\n", err)
	}
	err = cur.All(context.TODO(), &mhs)
	if err != nil {
		fmt.Println(err)
	}
	return mhs
}

func UpdateMhsById(db *mongo.Database, id primitive.ObjectID, npm int, nama string, fakultas model.Fakultas, dosen model.Dosen, programStudi model.ProgramStudi) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"npm":           npm,
			"nama":          nama,
			"fakultas":      fakultas,
			"program_studi": dosen,
			"dosen_wali":    programStudi,
		},
	}
	result, err := db.Collection("mahasiswa").UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateMahasiswa: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func DeleteMhsByID(db *mongo.Database, _id primitive.ObjectID) error {
	mhs := db.Collection("mahasiswa")
	filter := bson.M{"_id": _id}

	result, err := mhs.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// dosen
func InsertDosen(db *mongo.Database, kode string, nama string, hp string) (insertedID primitive.ObjectID, err error) {
	var dosen model.Dosen
	dosen.KodeDosen = kode
	dosen.Nama = nama
	dosen.PhoneNumber = hp
	result, err := db.Collection("dosen").InsertOne(context.Background(), dosen)
	if err != nil {
		fmt.Printf("InsertPresensi: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetDosenFromKodeDosen(db *mongo.Database, kode string) (dosen model.Dosen, errs error) {
	data_dhs := MongoConnect("db_dhs_tb").Collection("dosen")
	filter := bson.M{"kode_dosen": kode}
	err := data_dhs.FindOne(context.TODO(), filter).Decode(&dosen)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return dosen, fmt.Errorf("no data found for kode %s", kode)
		}
		return dosen, fmt.Errorf("error retrieving data for kode %s: %s", kode, err.Error())
	}
	return dosen, nil
}
func GetDosenFromID(db *mongo.Database, id primitive.ObjectID) (dosen model.Dosen, errs error) {
	data_dhs := MongoConnect("db_dhs_tb").Collection("dosen")
	filter := bson.M{"_id": id}
	err := data_dhs.FindOne(context.TODO(), filter).Decode(&dosen)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return dosen, fmt.Errorf("no data found for id %s", id)
		}
		return dosen, fmt.Errorf("error retrieving data for id %s: %s", id, err.Error())
	}
	return dosen, nil
}

func GetDosenAll(db *mongo.Database) (dosen []model.Dosen) {
	data_dosen := db.Collection("dosen")
	filter := bson.D{}
	// var results []mhs
	cur, err := data_dosen.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetMhsFromNPM: %v\n", err)
	}
	err = cur.All(context.TODO(), &dosen)
	if err != nil {
		fmt.Println(err)
	}
	return dosen
}

func UpdateDosenByID(db *mongo.Database, id primitive.ObjectID, kode string, nama string, hp string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"kode_dosen":   kode,
			"nama":         nama,
			"phone_number": hp,
		}}

	result, err := db.Collection("dosen").UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateDosen: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func DeleteDosenByID(db *mongo.Database, _id primitive.ObjectID) error {
	dosen := db.Collection("dosen")
	filter := bson.M{"_id": _id}

	result, err := dosen.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// matkul
func InsertMatkul(db *mongo.Database, kode string, nama string, sks int, dosen model.Dosen) (insertedID primitive.ObjectID, err error) {
	var matkul model.MataKuliah
	matkul.KodeMatkul = kode
	matkul.Nama = nama
	matkul.Sks = sks
	matkul.Dosen = dosen
	result, err := db.Collection("matkul").InsertOne(context.Background(), matkul)
	if err != nil {
		fmt.Printf("InsertMatkul: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetMatkulFromKodeMatkul(db *mongo.Database, kode string) (matkul model.MataKuliah, errs error) {
	data_matkul := db.Collection("matkul")
	filter := bson.M{"kode_matkul": kode}
	err := data_matkul.FindOne(context.TODO(), filter).Decode(&matkul)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return matkul, fmt.Errorf("no data found for kode %s", kode)
		}
		return matkul, fmt.Errorf("error retrieving data for kode %s: %s", kode, err.Error())
	}
	return matkul, nil
}

func GetMatkulFromID(db *mongo.Database, _id primitive.ObjectID) (matkul model.MataKuliah, errs error) {
	data_matkul := db.Collection("matkul")
	filter := bson.M{"_id": _id}
	err := data_matkul.FindOne(context.TODO(), filter).Decode(&matkul)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return matkul, fmt.Errorf("no data found for _id %s", _id)
		}
		return matkul, fmt.Errorf("error retrieving data for _id %s: %s", _id, err.Error())
	}
	return matkul, nil
}

func GetMatkulAll(db *mongo.Database) (matkul []model.MataKuliah) {
	data_matkul := db.Collection("matkul")
	filter := bson.D{}
	// var results []mhs
	cur, err := data_matkul.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetMhsFromMatkul: %v\n", err)
	}
	err = cur.All(context.TODO(), &matkul)
	if err != nil {
		fmt.Println(err)
	}
	return matkul
}

func UpdateMatkulFromID(db *mongo.Database, id primitive.ObjectID, kode string, nama string, sks int, dosen model.Dosen) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"kode_matkul": kode,
			"nama":        nama,
			"sks":         sks,
			"dosen":       dosen,
		}}

	result, err := db.Collection("matkul").UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateMatkul: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func DeleteMatkulByID(db *mongo.Database, _id primitive.ObjectID) error {
	dosen := db.Collection("matkul")
	filter := bson.M{"_id": _id}

	result, err := dosen.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// fakultas
func InsertFakultas(db *mongo.Database, kode string, nama string) (insertedID primitive.ObjectID, err error) {
	var fakultas model.Fakultas
	fakultas.KodeFakultas = kode
	fakultas.Nama = nama
	result, err := db.Collection("fakultas").InsertOne(context.Background(), fakultas)
	if err != nil {
		fmt.Printf("InsertFakultas: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetFakultasFromKodeFakultas(db *mongo.Database, kode string) (fakultas model.Fakultas, errs error) {
	data_fakultas := MongoConnect("db_dhs_tb").Collection("fakultas")
	filter := bson.M{"kode_fakultas": kode}
	err := data_fakultas.FindOne(context.TODO(), filter).Decode(&fakultas)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fakultas, fmt.Errorf("no data found for kode %s", kode)
		}
		return fakultas, fmt.Errorf("error retrieving data for kode %s: %s", kode, err.Error())
	}
	return fakultas, nil
}

func GetFakultasFromID(db *mongo.Database, _id primitive.ObjectID) (fakultas model.Fakultas, errs error) {
	data_fakultas := MongoConnect("db_dhs_tb").Collection("fakultas")
	filter := bson.M{"_id": _id}
	err := data_fakultas.FindOne(context.TODO(), filter).Decode(&fakultas)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fakultas, fmt.Errorf("no data found for _id %s", _id)
		}
		return fakultas, fmt.Errorf("error retrieving data for _Id %s: %s", _id, err.Error())
	}
	return fakultas, nil
}

func GetFakultasAll(db *mongo.Database) (fakultas []model.Fakultas) {
	data_mhs := MongoConnect("db_dhs_tb").Collection("fakultas")
	filter := bson.D{}
	// var results []mhs
	cur, err := data_mhs.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetMhsFromFakultas: %v\n", err)
	}
	err = cur.All(context.TODO(), &fakultas)
	if err != nil {
		fmt.Println(err)
	}
	return fakultas
}

func UpdateFakultasFromID(db *mongo.Database, id primitive.ObjectID, kode string, nama string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"kode_fakultas": kode,
			"nama":          nama,
		}}
	var fakultas model.Fakultas
	fakultas.KodeFakultas = kode
	fakultas.Nama = nama

	result, err := db.Collection("fakultas").UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateFakultas: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func DeleteFakultasFromID(db *mongo.Database, _id primitive.ObjectID) error {
	fakultas := db.Collection("fakultas")
	filter := bson.M{"_id": _id}

	result, err := fakultas.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// prodi
func InsertProdi(db *mongo.Database, kode string, nama string) (insertedID primitive.ObjectID, err error) {
	var programStudi model.ProgramStudi
	programStudi.KodeProgramStudi = kode
	programStudi.Nama = nama
	result, err := db.Collection("programStudi").InsertOne(context.Background(), programStudi)
	if err != nil {
		fmt.Printf("InsertProdi: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetProdiFromKodeProdi(db *mongo.Database, kode string) (programStudi model.ProgramStudi, errs error) {
	data_dhs := MongoConnect("db_dhs_tb").Collection("programStudi")
	filter := bson.M{"kode_program_studi": kode}
	err := data_dhs.FindOne(context.TODO(), filter).Decode(&programStudi)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return programStudi, fmt.Errorf("no data found for kode %s", kode)
		}
		return programStudi, fmt.Errorf("error retrieving data for kode %s: %s", kode, err.Error())
	}
	// if err != nil {
	// 	fmt.Printf("GetProdiFromKodeProdi: %v\n", err)
	// }
	return programStudi, nil
}

func GetProdiFromID(db *mongo.Database, _id primitive.ObjectID) (prodi model.ProgramStudi, errs error) {
	data_prodi := MongoConnect("db_dhs_tb").Collection("programStudi")
	filter := bson.M{"_id": _id}
	err := data_prodi.FindOne(context.TODO(), filter).Decode(&prodi)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return prodi, fmt.Errorf("no data found for _id %s", _id)
		}
		return prodi, fmt.Errorf("error retrieving data for _Id %s: %s", _id, err.Error())
	}
	return prodi, nil
}

func GetProdiAll(db *mongo.Database) (programStudi []model.ProgramStudi) {
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

func UpdateProdiByID(db *mongo.Database, id primitive.ObjectID, kode string, nama string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"kode_program_studi": kode,
			"nama":               nama,
		}}

	result, err := db.Collection("programStudi").UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateProdi: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func DeleteProdiAll(db *mongo.Database, _id primitive.ObjectID) error {
	programStudi := db.Collection("programStudi")
	filter := bson.M{"_id": _id}

	result, err := programStudi.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
