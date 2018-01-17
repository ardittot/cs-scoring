package main

import (
    // "gopkg.in/resty.v1"
)

// Functions Declaration
/*
func InitKafkaConsumer() {
    resty.R().SetBody("{\"topic\":\"crs-unscored\",\"group\":\"test-group\"}").Post("http://localhost:8020/subscribe/topic/add")
    resty.R().SetBody("{\"data\":[{\"topic\":\"crs-unscored\",\"url\":[\"http://0.0.0.0:8000/crs\"]}]}").Post("http://localhost:8020/subscribe/url/add")
}

func ProduceKafka(data interface{}) {
    resty.R().SetBody(data).Post("http://localhost:8020/publish/crs-scored")
}
*/

// Variable Types Declaration

// Las_t_kupedes_scoring
type Las_t_kupedes_scoring struct {
    ID_Scoring uint64 `json:"id_scoring"`
    Net_income uint64 `json:"net_income"`
    Angsuran uint64 `json:"angsuran"`
    Nilai_likuidasi_agunan uint64 `json:"nilai_likuidasi_agunan"`
    Plafon uint64 `json:"plafon"`
    Jenis_bukti_kepemilikan_agunan string `json:"jenis_bukti_kepemilikan_agunan"`
    Usia uint `json:"usia"`
    Debitur_baru string `json:"debitur_baru"`
    Lama_usaha uint `json:"lama_usaha"`
    Status_perkawinan string `json:"status_perkawinan"`
    Punya_usaha_sampingan string `json:"punya_usaha_sampingan"`
    Tujuan_penggunaan_kredit string `json:"tujuan_penggunaan_kredit"`
    Punya_pelanggan_tetap string `json:"punya_pelanggan_tetap"`
}

func (l Las_t_kupedes_scoring) ToClean() Las_t_kupedes_scoring_clean {
    var l1 Las_t_kupedes_scoring_clean
    l1.ID_Scoring              = l.ID_Scoring
    // Kriteria 1
    l1.Net_income              = l.Net_income
    l1.Angsuran                = l.Angsuran
    // Kriteria 2
    l1.Nilai_likuidasi_agunan  = l.Nilai_likuidasi_agunan
    l1.Plafon                  = l.Plafon
    // Kriteria 3
    l1.Jenis_bukti_kepemilikan_agunan_Kwitansi = 0
    l1.Jenis_bukti_kepemilikan_agunan_BPKB = 0
    l1.Jenis_bukti_kepemilikan_agunan_Sertifikat = 0
    l1.Jenis_bukti_kepemilikan_agunan_SuratTanahDiluarSertifikat = 0
    l1.Jenis_bukti_kepemilikan_agunan_PetokD = 0
    l1.Jenis_bukti_kepemilikan_agunan_LetterC = 0
    switch {
    case l.Jenis_bukti_kepemilikan_agunan == "Kwitansi":
        l1.Jenis_bukti_kepemilikan_agunan_Kwitansi  = 1
    case l.Jenis_bukti_kepemilikan_agunan == "BPKB":
        l1.Jenis_bukti_kepemilikan_agunan_BPKB = 1
    case l.Jenis_bukti_kepemilikan_agunan == "Sertifikat":
        l1.Jenis_bukti_kepemilikan_agunan_Sertifikat = 1
    case l.Jenis_bukti_kepemilikan_agunan == "Surat Tanah Diluar Sertifikat":
        l1.Jenis_bukti_kepemilikan_agunan_SuratTanahDiluarSertifikat = 1
    case l.Jenis_bukti_kepemilikan_agunan == "Petok D":
        l1.Jenis_bukti_kepemilikan_agunan_PetokD = 1
    case l.Jenis_bukti_kepemilikan_agunan == "Letter C":
        l1.Jenis_bukti_kepemilikan_agunan_LetterC = 1
    }
    // Kriteria 4
    l1.Usia                    = l.Usia
    // Kriteria 5
    l1.Debitur_lama  = 0
    if l.Debitur_baru == "Ya" {
        l1.Debitur_lama  = 1
    }
    // Kriteria 6
    l1.Lama_usaha              = l.Lama_usaha
    // Kriteria 7
    l1.Status_perkawinan_Menikah = 0
    l1.Status_perkawinan_BelumMenikah = 0
    l1.Status_perkawinan_DudaJanda = 0
    switch {
    case l.Status_perkawinan == "Menikah":
        l1.Status_perkawinan_Menikah  = 1
    case l.Status_perkawinan == "Belum Menikah":
        l1.Status_perkawinan_BelumMenikah = 1
    case l.Status_perkawinan == "Duda/Janda":
        l1.Status_perkawinan_DudaJanda = 1
    }
    // Kriteria 8
    l1.Punya_usaha_sampingan_Ya  = 0
    if l.Punya_usaha_sampingan == "Ya" {
        l1.Punya_usaha_sampingan_Ya  = 1
    }
    // Kriteria 9
    l1.Tujuan_penggunaan_kredit_ModalKerja = 0
    l1.Tujuan_penggunaan_kredit_PenggantiModalKerja = 0
    l1.Tujuan_penggunaan_kredit_Investasi = 0
    switch {
    case l.Tujuan_penggunaan_kredit == "Modal Kerja":
        l1.Tujuan_penggunaan_kredit_ModalKerja  = 1
    case l.Tujuan_penggunaan_kredit == "Pengganti Modal Kerja":
        l1.Tujuan_penggunaan_kredit_PenggantiModalKerja = 1
    case l.Tujuan_penggunaan_kredit == "Investasi":
        l1.Tujuan_penggunaan_kredit_Investasi = 1
    }
    // Kriteria 10
    l1.Punya_pelanggan_tetap_Ya  = 0
    if l.Punya_pelanggan_tetap == "Ya" {
        l1.Punya_pelanggan_tetap_Ya  = 1
    }

    return l1
}

// Las_t_kupedes_scoring_CLEAN
type Las_t_kupedes_scoring_clean struct {
    ID_Scoring uint64
    Net_income uint64
    Angsuran uint64
    Nilai_likuidasi_agunan uint64
    Plafon uint64
    Jenis_bukti_kepemilikan_agunan_Kwitansi uint
    Jenis_bukti_kepemilikan_agunan_SuratTanahDiluarSertifikat uint
    Jenis_bukti_kepemilikan_agunan_PetokD uint
    Jenis_bukti_kepemilikan_agunan_LetterC uint
    Jenis_bukti_kepemilikan_agunan_BPKB uint
    Jenis_bukti_kepemilikan_agunan_Sertifikat uint
    Usia uint
    Debitur_lama uint
    Lama_usaha uint
    Status_perkawinan_Menikah uint
    Status_perkawinan_BelumMenikah uint
    Status_perkawinan_DudaJanda uint
    Punya_usaha_sampingan_Ya uint
    Tujuan_penggunaan_kredit_ModalKerja uint
    Tujuan_penggunaan_kredit_PenggantiModalKerja uint
    Tujuan_penggunaan_kredit_Investasi uint
    Punya_pelanggan_tetap_Ya uint
}

func (l Las_t_kupedes_scoring_clean) Score() Las_kupedes_scored {
    var l1 Las_kupedes_scored
    var score uint
    var grade uint
    score = 0
    l1.ID_Scoring   = l.ID_Scoring
    // Kriteria 1
    ratio_income_to_angsuran := float32(l.Net_income) / float32(l.Angsuran)
    switch {
    case ratio_income_to_angsuran > 3:
        score += 325
    case ratio_income_to_angsuran >= 2:
        score += 290
    case ratio_income_to_angsuran >= 1.33:
        score += 225
    }
    // Kriteria 2
    ratio_agunan_to_plafon := float32(l.Nilai_likuidasi_agunan) / float32(l.Plafon)
    switch {
    case ratio_agunan_to_plafon > 2:
        score += 155
    case ratio_agunan_to_plafon > 1.5:
        score += 125
    case ratio_agunan_to_plafon >= 1:
        score += 65
    }
    // Kriteria 3
    score += l.Jenis_bukti_kepemilikan_agunan_Sertifikat * 160
    score += l.Jenis_bukti_kepemilikan_agunan_BPKB * 115
    score += l.Jenis_bukti_kepemilikan_agunan_PetokD * 55
    score += l.Jenis_bukti_kepemilikan_agunan_LetterC * 55
    score += l.Jenis_bukti_kepemilikan_agunan_SuratTanahDiluarSertifikat * 55
    // Kriteria 4
    switch {
    case l.Usia > 50:
        score += 40
    case l.Usia > 45:
        score += 30
    case l.Usia >= 28:
        score += 5
    }
    // Kriteria 5
    if l.Debitur_lama > 0 {
        score += 20
    }
    // Kriteria 6
    switch {
    case l.Lama_usaha > 8:
        score += 80
    case l.Lama_usaha > 5:
        score += 70
    case l.Lama_usaha >= 2:
        score += 45
    }
    // Kriteria 7
    score += l.Status_perkawinan_BelumMenikah * 115
    score += l.Status_perkawinan_Menikah * 100
    // Kriteria 8
    if l.Punya_usaha_sampingan_Ya > 0 {
        score += 45
    }
    // Kriteria 9
    score += l.Tujuan_penggunaan_kredit_Investasi * 70
    score += l.Tujuan_penggunaan_kredit_PenggantiModalKerja * 15
    // Kriteria 10
    if l.Punya_pelanggan_tetap_Ya > 0 {
        score += 60
    }
    // Score
    l1.Score = float32(score)
    // Grade
    switch {
    case l1.Score > 950:
        grade = 1
    case l1.Score > 875:
        grade = 2
    case l1.Score > 800:
        grade = 3
    case l1.Score > 725:
        grade = 4
    case l1.Score > 650:
        grade = 5
    case l1.Score > 575:
        grade = 6
    case l1.Score > 500:
        grade = 7
    case l1.Score > 425:
        grade = 8
    case l1.Score > 350:
        grade = 9
    default:
        grade = 10
    }
    l1.Grade = uint8(grade)
    return l1
}

// LAS_STATUS
type Las_kupedes_scored struct {
    ID_Scoring uint64 `json:"id_scoring" default:nil`
    Score float32 `json:"score" default:nil`
    Grade uint8 `json:"grade" default:nil`
}
