package namapackage

import (
	"fmt"
	"testing"

	model "github.com/daniferdinandall/be_dhs_p3/model"
	module "github.com/daniferdinandall/be_dhs_p3/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

// ========================
// DHS
// dhs
func TestInsertDhs(t *testing.T) {
	mahasiswa := model.Mahasiswa{
		Npm:  1214030,
		Nama: "Aryka Anisa Pertiwi",
		Fakultas: model.Fakultas{
			KodeFakultas: "SV",
			Nama:         "Sekolah Vokasi",
		},
		ProgramStudi: model.ProgramStudi{
			KodeProgramStudi: "D4TI",
			Nama:             "D4 Teknik Informatika",
		},
		DosenWali: model.Dosen{
			PhoneNumber: "086653219827",
			KodeDosen:   "NSF",
			Nama:        "Rd. NURAINI SITI FATHONAH, S.S., M.Hum.,SFPC",
		},
	}
	mata_kuliah := []model.NilaiMataKuliah{
		{
			KodeMatkul: "TI41061",
			Nama:       "ALGORITMA DAN STRUKTUR DATA I",
			Sks:        3,
			Nilai:      "A",
		}, {
			KodeMatkul: "TI41092",
			Nama:       "ALJABAR LINIER",
			Sks:        2,
			Nilai:      "A",
		}, {
			KodeMatkul: "PPI01040",
			Nama:       "BAHASA INDONESIA",
			Sks:        2,
			Nilai:      "AB",
		}, {
			KodeMatkul: "TI42011",
			Nama:       "LITERASI TEKNOLOGI",
			Sks:        2,
			Nilai:      "AB",
		}, {
			KodeMatkul: "TI41071",
			Nama:       "ALGORITMAPEMOGRAMAN I",
			Sks:        3,
			Nilai:      "A",
		},
	}

	insertedID, err := module.InsertDHS(module.MongoConn, mahasiswa, mata_kuliah)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}
func TestGetDhsFromNPM(t *testing.T) {
	npm := 1214070
	biodata, err := module.GetDhsFromNPM(module.MongoConn, npm)
	// biodata, err := module.GetDhsFromID(module.MongoConn, id)
	if err != nil {
		fmt.Println("erorr")
	} else {
		fmt.Println(biodata)
	}
}
func TestGetDhsFromID(t *testing.T) {
	id, er := primitive.ObjectIDFromHex("648c20a002e751343198b97a")
	if er != nil {
		fmt.Println("Invalid id parameter")
	}
	biodata, err := module.GetDhsFromID(module.MongoConn, id)
	if err != nil {
		fmt.Println("erorr")
	} else {
		fmt.Println(biodata)
	}
}
func TestGetDhsAll(t *testing.T) {
	biodata := module.GetDhsAll(module.MongoConn)
	fmt.Println(biodata)
}
func TestEditDhs(t *testing.T) {
	mahasiswa := model.Mahasiswa{
		Npm:  1214030,
		Nama: "Aryka Anisa Pertiwi second",
		Fakultas: model.Fakultas{
			KodeFakultas: "SV",
			Nama:         "Sekolah Vokasi",
		},
		ProgramStudi: model.ProgramStudi{
			KodeProgramStudi: "D4TI",
			Nama:             "D4 Teknik Informatika",
		},
		DosenWali: model.Dosen{
			PhoneNumber: "086653219827",
			KodeDosen:   "NSF",
			Nama:        "Rd. NURAINI SITI FATHONAH, S.S., M.Hum.,SFPC",
		},
	}
	mata_kuliah := []model.NilaiMataKuliah{
		{
			KodeMatkul: "TI41061",
			Nama:       "ALGORITMA DAN STRUKTUR DATA I",
			Sks:        3,
			Nilai:      "A",
		}, {
			KodeMatkul: "TI41092",
			Nama:       "ALJABAR LINIER",
			Sks:        2,
			Nilai:      "A",
		}, {
			KodeMatkul: "PPI01040",
			Nama:       "BAHASA INDONESIA",
			Sks:        2,
			Nilai:      "AB",
		}, {
			KodeMatkul: "TI42011",
			Nama:       "LITERASI TEKNOLOGI",
			Sks:        2,
			Nilai:      "AB",
		}, {
			KodeMatkul: "TI41071",
			Nama:       "ALGORITMAPEMOGRAMAN I",
			Sks:        3,
			Nilai:      "A",
		},
	}
	id, err := primitive.ObjectIDFromHex("648c20a002e751343198b97a")
	if err != nil {
		fmt.Printf("Data tidak berhasil disimpan dengan id")
	} else {
		module.UpdateDhsById(module.MongoConn, id, mahasiswa, mata_kuliah)
		fmt.Printf("Data berhasil disimpan dengan id")
	}
}
func TestDeleteDhs(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("648c20a002e751343198b97a")
	if err != nil {
		fmt.Printf("Data tidak berhasil disimpan dengan id")
	} else {
		biodata := module.DeleteDhsByID(module.MongoConn, id)
		fmt.Println(biodata)
	}
}

// mhs
func TestInsertMhs(t *testing.T) {
	npm := 1214031
	nama := "Erdito Nausah Adam"
	fakultas := model.Fakultas{
		KodeFakultas: "SV",
		Nama:         "Sekolah Vokasi",
	}
	programStudi := model.ProgramStudi{
		KodeProgramStudi: "D4TI",
		Nama:             "D4 Teknik Informatika",
	}
	dosen := model.Dosen{
		KodeDosen:   "NSF",
		PhoneNumber: "086653219827",
		Nama:        " Rd. NURAINI SITI FATHONAH, S.S., M.Hum.,SFPC",
	}

	hasil, err := module.InsertMhs(module.MongoConn, npm, nama, fakultas, dosen, programStudi)
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", hasil.Hex())
}

func TestGetMhsFromNPM(t *testing.T) {
	npm := 1214031
	biodata, err := module.GetMhsFromNPM(module.MongoConn, npm)
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetMhsID(t *testing.T) {
	// npm := 1214070
	id, er := primitive.ObjectIDFromHex("648c27037b1ba3b82b383d06")
	if er != nil {
		fmt.Println("Invalid id parameter")
	}
	biodata, err := module.GetMhsFromID(module.MongoConn, id)
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetMhsAll(t *testing.T) {
	biodata := module.GetMhsAll(module.MongoConn)
	fmt.Println(biodata)
}
func TestEditMhs(t *testing.T) {
	npm := 1214031
	nama := "Erdito Nausah Hawa"
	fakultas := model.Fakultas{
		KodeFakultas: "SV",
		Nama:         "Sekolah Vokasi",
	}
	programStudi := model.ProgramStudi{
		KodeProgramStudi: "D4TI",
		Nama:             "D4 Teknik Informatika",
	}
	dosen := model.Dosen{
		KodeDosen:   "NSF",
		PhoneNumber: "086653219827",
		Nama:        "Rd. NURAINI SITI FATHONAH, S.S., M.Hum.,SFPC",
	}

	id, err := primitive.ObjectIDFromHex("648c27037b1ba3b82b383d06")
	if err != nil {
		fmt.Printf("Data tidak berhasil disimpan dengan id")
	} else {
		module.UpdateMhsById(module.MongoConn, id, npm, nama, fakultas, dosen, programStudi)
		fmt.Printf("Data berhasil disimpan dengan id")
	}
}
func TestDeleteMhs(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("648c27037b1ba3b82b383d06")
	if err != nil {
		fmt.Printf("Data tidak berhasil disimpan dengan id")
	} else {
		biodata := module.DeleteMhsByID(module.MongoConn, id)
		fmt.Println(biodata)
	}
}

// dosen
func TestInsertDosen(t *testing.T) {
	PhoneNumber := "086653219827"
	KodeDosen := "NSF"
	Nama := "Rd. NURAINI SITI FATHONAH, S.S., M.Hum.,SFPC"

	hasil, err := module.InsertDosen(module.MongoConn, KodeDosen, Nama, PhoneNumber)
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	fmt.Println(hasil)
}

func TestGetDosenFromKodeDosen(t *testing.T) {
	kode := "MYS"
	biodata, err := module.GetDosenFromKodeDosen(module.MongoConn, kode)
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	fmt.Println(biodata)
}
func TestGetDosenFromId(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("648c6eefc145cbe151f52c85")
	biodata, err := module.GetDosenFromID(module.MongoConn, id)
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetDosenAll(t *testing.T) {
	biodata := module.GetDosenAll(module.MongoConn)
	fmt.Println(biodata)
}

func TestEditDosen(t *testing.T) {
	PhoneNumber := "086653219827"
	KodeDosen := "NSF"
	Nama := "Rd. NURAINI s FATHONAH, S.S., M.Hum.,SFPC"
	id, err := primitive.ObjectIDFromHex("648c6eefc145cbe151f52c85")
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	hasil := module.UpdateDosenByID(module.MongoConn, id, KodeDosen, Nama, PhoneNumber)
	fmt.Println(hasil)
}
func TestDeleteDosen(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("648c6eefc145cbe151f52c85")
	if err != nil {
		fmt.Printf("Data tidak berhasil disimpan dengan id")
	} else {
		biodata := module.DeleteDosenByID(module.MongoConn, id)
		fmt.Println(biodata)
	}
}

// dosen

// MATKUL
func TestInsertMatkul(t *testing.T) {
	kode := "TI41264"
	nama := "PEMOGRAMAN III (WEBSERVICE)"
	sks := 3
	dosen := model.Dosen{
		Nama: "Indra Riksa Herlambang",
	}

	hasil, err := module.InsertMatkul(module.MongoConn, kode, nama, sks, dosen)
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	fmt.Println(hasil)
}

func TestMatkulFromKodeMatkul(t *testing.T) {
	kode := "TI41264"
	biodata, err := module.GetMatkulFromKodeMatkul(module.MongoConn, kode)
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	fmt.Println(biodata)
}
func TestMatkulFromID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("648c719bf744e25b91346427")
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	biodata, err := module.GetMatkulFromID(module.MongoConn, id)
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	fmt.Println(biodata)
}

func TestMatkulAll(t *testing.T) {
	biodata := module.GetMatkulAll(module.MongoConn)
	fmt.Println(biodata)
}

func TestUpdateMatkul(t *testing.T) {
	kode := "TI41264"
	nama := "PEMOGRAMAN III2 (WEBSERVICE)"
	sks := 3
	dosen := model.Dosen{
		Nama: "Indra Riksa Herlambang",
	}
	id, err := primitive.ObjectIDFromHex("648c719bf744e25b91346427")
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	hasil := module.UpdateMatkulFromID(module.MongoConn, id, kode, nama, sks, dosen)
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	fmt.Println(hasil)
}

func TestDeleteMatkul(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("648c719bf744e25b91346427")
	if err != nil {
		fmt.Printf("Data tidak berhasil disimpan dengan id")
	} else {
		biodata := module.DeleteMatkulByID(module.MongoConn, id)
		fmt.Println(biodata)
	}
}

// FAKULTAS
func TestInsertFakultas(t *testing.T) {
	kode := "SV"
	nama := "Sekolah Vokasi"

	hasil, err := module.InsertFakultas(module.MongoConn, kode, nama)
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	fmt.Println(hasil)
}

func TestFakultasFromKodeFakultas(t *testing.T) {
	kode := "SV"
	biodata, err := module.GetFakultasFromKodeFakultas(module.MongoConn, kode)
	if err != nil {
		fmt.Printf("Error get data: %v", err)
	}
	fmt.Println(biodata)
}
func TestFakultasFromID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("648c740e9fd2ec61fc370d60")
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	biodata, err := module.GetFakultasFromID(module.MongoConn, id)
	if err != nil {
		fmt.Printf("Error get data: %v", err)
	}
	fmt.Println(biodata)
}

func TestFakultasAll(t *testing.T) {
	biodata := module.GetFakultasAll(module.MongoConn)
	fmt.Println(biodata)
}
func TestUpdateFakultas(t *testing.T) {
	kode := "SVq"
	nama := "Sekolah Vokasi"
	id, err := primitive.ObjectIDFromHex("648c740e9fd2ec61fc370d60")
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	hasil := module.UpdateFakultasFromID(module.MongoConn, id, kode, nama)
	fmt.Println(hasil)
}

func TestDeleteFakultas(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("648c740e9fd2ec61fc370d60")
	if err != nil {
		fmt.Printf("Data tidak berhasil disimpan dengan id")
	} else {
		biodata := module.DeleteFakultasFromID(module.MongoConn, id)
		fmt.Println(biodata)
	}
}

// Prodi
func TestInsertProdi(t *testing.T) {
	kode := "D4TI"
	nama := "D4 Teknik Informatika"

	hasil, err := module.InsertProdi(module.MongoConn, kode, nama)
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	fmt.Println(hasil)
}

func TestProdiFromKodeProdi(t *testing.T) {
	kode := "D4TI"
	biodata, err := module.GetProdiFromKodeProdi(module.MongoConn, kode)
	if err != nil {
		fmt.Printf("Error get data: %v", err)
	}
	fmt.Println(biodata)
}
func TestProdiFromID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("648c770932bf809c88fe5222")
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	biodata, err := module.GetProdiFromID(module.MongoConn, id)
	if err != nil {
		fmt.Printf("Error get data: %v", err)
	}
	fmt.Println(biodata)
}

func TestProdiAll(t *testing.T) {
	biodata := module.GetProdiAll(module.MongoConn)
	fmt.Println(biodata)
}

func TestEditProdi(t *testing.T) {
	kode := "D4TIw"
	nama := "D4s Teknik Informatika"
	id, err := primitive.ObjectIDFromHex("648c770932bf809c88fe5222")
	if err != nil {
		fmt.Printf("Error inserting data: %v", err)
	}
	hasil := module.UpdateProdiByID(module.MongoConn, id, kode, nama)
	fmt.Println(hasil)
}
func TestDeleteProdi(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("648c770932bf809c88fe5222")
	if err != nil {
		fmt.Printf("Data tidak berhasil disimpan dengan id")
	} else {
		biodata := module.DeleteProdiByID(module.MongoConn, id)
		fmt.Println(biodata)
	}
}
