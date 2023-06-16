let Mahasiswa = [
    {
        nama,
        npm,
        Fakultas,
        Program_studi,
        dosen_wali: Dosen
    }
]

let Dosen = [
    {
        kode_dosen:"string",
        nama_dosen:"string"
    }
]

let Matkul = [
    {
        kode_matkul:"string",
        nama_matkul:"string",
        sks:4,
        dosen: Dosen[0]
    }
]

let Fakultas = [{
    kode_fakultas,
    nama_fakultasi
}]

let Program_studi = [{
    kode_fakultas,
    nama_fakultasi
}]

let Dhs = [
    {
        mahasiswa:Mahasiswa,
        mata_kuliah: [
            Matkul
        ]
    }
]