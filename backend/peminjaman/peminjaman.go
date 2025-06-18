package peminjaman

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Peminjaman struct {
	ID                 int    `json:"id"`
	BarangID           int    `json:"barang_id"`
	NamaBarang         string `json:"nama_barang"`
	Peminjam           string `json:"peminjam"`
	TanggalPeminjam    string `json:"tanggal_peminjam"`
	TanggalKembali     string `json:"tanggal_kembali"`
	StatusPengembalian string `json:"status_pengembalian"`
}

func connectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/inventaris_barang")
}

func GetAllPeminjaman(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB connection error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT 
		peminjaman.id, 
		peminjaman.barang_id, 
		barang.nama AS nama_barang, 
		peminjaman.peminjam, 
		peminjaman.tanggal_peminjam, 
		peminjaman.tanggal_kembali, 
		peminjaman.status_pengembalian
	FROM peminjaman
	JOIN barang ON peminjaman.barang_id = barang.id`)

	if err != nil {
		http.Error(w, "Query error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var list []Peminjaman
	for rows.Next() {
		var p Peminjaman
		if err := rows.Scan(&p.ID, &p.BarangID, &p.NamaBarang, &p.Peminjam, &p.TanggalPeminjam, &p.TanggalKembali, &p.StatusPengembalian); err != nil {
			http.Error(w, "Scan error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		list = append(list, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func TambahPeminjaman(w http.ResponseWriter, r *http.Request) {
	var data Peminjaman
	json.NewDecoder(r.Body).Decode(&data)

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(`
		INSERT INTO peminjaman (barang_id, peminjam, tanggal_peminjam, tanggal_kembali, status_pengembalian)
		VALUES (?, ?, ?, ?, ?)`,
		data.BarangID, data.Peminjam, data.TanggalPeminjam, data.TanggalKembali, data.StatusPengembalian)

	if err != nil {
		http.Error(w, "Gagal tambah data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("UPDATE barang SET status = 'Dipinjam' WHERE id = ?", data.BarangID)
	if err != nil {
		http.Error(w, "Gagal update status barang", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func EditPeminjaman(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	var data Peminjaman
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Data tidak valid: "+err.Error(), http.StatusBadRequest)
		return
	}

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(`
		UPDATE peminjaman
		SET barang_id = ?, peminjam = ?, tanggal_peminjam = ?, tanggal_kembali = ?, status_pengembalian = ?
		WHERE id = ?`,
		data.BarangID, data.Peminjam, data.TanggalPeminjam, data.TanggalKembali, data.StatusPengembalian, id)

	if err != nil {
		http.Error(w, "Gagal update: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//  Jika dikembalikan, update status barang ke "Tersedia"
	if data.StatusPengembalian == "Dikembalikan" {
		_, err = db.Exec("UPDATE barang SET status = 'Tersedia' WHERE id = ?", data.BarangID)
		if err != nil {
			http.Error(w, "Gagal update status barang", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Data peminjaman berhasil diupdate",
	})
}

func HapusPeminjaman(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Ambil barang_id dari peminjaman yang akan dihapus
	var barangID int
	err = db.QueryRow("SELECT barang_id FROM peminjaman WHERE id = ?", id).Scan(&barangID)
	if err != nil {
		http.Error(w, "Data peminjaman tidak ditemukan", http.StatusNotFound)
		return
	}

	// Update status barang menjadi "Tersedia"
	_, err = db.Exec("UPDATE barang SET status = 'Tersedia' WHERE id = ?", barangID)
	if err != nil {
		http.Error(w, "Gagal update status barang", http.StatusInternalServerError)
		return
	}

	// Hapus data peminjaman
	result, err := db.Exec("DELETE FROM peminjaman WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Gagal hapus data peminjaman", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "ID tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Data peminjaman berhasil dihapus dan status barang diperbarui",
	})
}

func KembalikanBarang(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var barangID int
	err = db.QueryRow("SELECT barang_id FROM peminjaman WHERE id = ?", id).Scan(&barangID)
	if err != nil {
		http.Error(w, "Gagal ambil barang_id", http.StatusInternalServerError)
		return
	}

	// Ubah status peminjaman & barang
	_, _ = db.Exec("UPDATE peminjaman SET status_pengembalian = 'Dikembalikan' WHERE id = ?", id)
	_, _ = db.Exec("UPDATE barang SET status = 'Tersedia' WHERE id = ?", barangID)

	json.NewEncoder(w).Encode(map[string]string{"message": "Barang berhasil dikembalikan"})
}
