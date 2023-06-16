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

// // mhs
// func TestInsertMhs(t *testing.T) {
// 	npm := 1214070
// 	nama := "Farel Nouval Daswara"
// 	fakultas := model.Fakultas{
// 		KodeFakultas: "SV",
// 		Nama:         "Sekolah Vokasi",
// 	}
// 	programStudi := model.ProgramStudi{
// 		KodeProgramStudi: "D4TI",
// 		Nama:             "D4 Teknik Informatika",
// 	}
// 	dosen := model.Dosen{
// 		KodeDosen:   "NSF",
// 		PhoneNumber: "086653219827",
// 		Nama:        " Rd. NURAINI SITI FATHONAH, S.S., M.Hum.,SFPC",
// 	}

// 	hasil := module.InsertMhs(npm, nama, fakultas, dosen, programStudi)
// 	fmt.Println(hasil)
// }

// func TestGetMhsFromNPM(t *testing.T) {
// 	npm := 1214070
// 	biodata := module.GetMhsFromNPM(npm)
// 	fmt.Println(biodata)
// }

// func TestGetMhsAll(t *testing.T) {
// 	biodata := module.GetMhsAll()
// 	fmt.Println(biodata)
// }

// // dosen
// func TestInsertDosen(t *testing.T) {
// 	PhoneNumber := "086653219827"
// 	KodeDosen := "NSF"
// 	Nama := "Rd. NURAINI SITI FATHONAH, S.S., M.Hum.,SFPC"

// 	hasil := module.InsertDosen(KodeDosen, Nama, PhoneNumber)
// 	fmt.Println(hasil)
// }

// func TestGetDosenFromKodeDosen(t *testing.T) {
// 	kode := "NSF"
// 	biodata := module.GetDosenFromKodeDosen(kode)
// 	fmt.Println(biodata)
// }

// func TestGetDosenAll(t *testing.T) {
// 	biodata := module.GetDosenAll()
// 	fmt.Println(biodata)
// }

// // dosen

// // MATKUL
// func TestInsertMatkul(t *testing.T) {
// 	kode := "TI41264"
// 	nama := "PEMOGRAMAN III (WEBSERVICE)"
// 	sks := 3
// 	dosen := model.Dosen{
// 		Nama: "Indra Riksa Herlambang",
// 	}

// 	hasil := module.InsertMatkul(kode, nama, sks, dosen)
// 	fmt.Println(hasil)
// }

// func TestMatkulFromKodeMatkul(t *testing.T) {
// 	kode := "TI41264"
// 	biodata := module.GetMatkulFromKodeMatkul(kode)
// 	fmt.Println(biodata)
// }

// func TestMatkulAll(t *testing.T) {
// 	biodata := module.GetMatkulAll()
// 	fmt.Println(biodata)
// }

// // FAKULTAS
// func TestInsertFakultas(t *testing.T) {
// 	kode := "SV"
// 	nama := "Sekolah Vokasi"

// 	hasil := module.InsertFakultas(kode, nama)
// 	fmt.Println(hasil)
// }

// func TestFakultasFromKodeFakultas(t *testing.T) {
// 	kode := "SV"
// 	biodata := module.GetFakultasFromKodeFakultas(kode)
// 	fmt.Println(biodata)
// }

// func TestFakultasAll(t *testing.T) {
// 	biodata := module.GetFakultasAll()
// 	fmt.Println(biodata)
// }

// // Prodi
// func TestInsertProdi(t *testing.T) {
// 	kode := "D4TI"
// 	nama := "D4 Teknik Informatika"

// 	hasil := module.InsertProdi(kode, nama)
// 	fmt.Println(hasil)
// }

// func TestProdiFromKodeProdi(t *testing.T) {
// 	kode := "D4TI"
// 	biodata := module.GetProdiFromKodeProdi(kode)
// 	fmt.Println(biodata)
// }

// func TestProdiAll(t *testing.T) {
// 	biodata := module.GetProdiAll()
// 	fmt.Println(biodata)
// }
