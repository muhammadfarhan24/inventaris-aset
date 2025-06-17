package dashboard

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type DashboardSummary struct {
	TotalBarang   int `json:"total_barang"`
	TotalDipinjam int `json:"total_dipinjam"`
	TotalService  int `json:"total_service"`
	TotalPengguna int `json:"total_pengguna"`
}

func GetSummaryDashboard(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/inventaris_barang")
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var summary DashboardSummary

	db.QueryRow("SELECT COUNT(*) FROM barang").Scan(&summary.TotalBarang)
	db.QueryRow("SELECT COUNT(*) FROM peminjaman WHERE status_pengembalian = 'Dipinjam'").Scan(&summary.TotalDipinjam)
	db.QueryRow("SELECT COUNT(*) FROM service_barang WHERE status = 'Dalam Service'").Scan(&summary.TotalService)
	db.QueryRow("SELECT COUNT(*) FROM user").Scan(&summary.TotalPengguna)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}
