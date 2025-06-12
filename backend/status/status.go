package status

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type StatusSummary struct {
	Status string `json:"status"`
	Jumlah int    `json:"jumlah"`
}

type Barang struct {
	ID     int    `json:"id"`
	Nama   string `json:"nama"`
	Status string `json:"status"`
}

func connectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/inventaris_barang")
}

func GetStatusSummary(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT status, COUNT(*) as jumlah FROM barang GROUP BY status`)
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var summary []StatusSummary
	for rows.Next() {
		var s StatusSummary
		rows.Scan(&s.Status, &s.Jumlah)
		summary = append(summary, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

func GetBarangByStatus(w http.ResponseWriter, r *http.Request) {
	statusParam := r.URL.Query().Get("status")
	if statusParam == "" {
		http.Error(w, "Parameter status kosong", http.StatusBadRequest)
		return
	}

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, nama, status FROM barang WHERE status = ?`, statusParam)
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var barangList []Barang
	for rows.Next() {
		var b Barang
		rows.Scan(&b.ID, &b.Nama, &b.Status)
		barangList = append(barangList, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(barangList)
}
