<template>
  <div class="main-content">
    <header class="header">
      <h1>Daftar Barang</h1>
    </header>

    <section class="content">
      <h2>Ringkasan Status Barang</h2>

      <div class="cards">
        <div class="card available" @click="setStatus('Tersedia')">
          <h3>Barang Tersedia</h3>
          <p>80</p>
        </div>
        <div class="card damaged" @click="setStatus('Rusak')">
          <h3>Barang Rusak</h3>
          <p>25</p>
        </div>
        <div class="card borrowed" @click="setStatus('Dipinjam')">
          <h3>Barang Dipinjam</h3>
          <p>15</p>
        </div>
      </div>

      <!-- Tabel berdasarkan status yang diklik -->
      <div v-if="activeStatus" class="table-section">
        <button class="close-btn" @click="activeStatus = ''">Tutup</button>
        <h3>Daftar Barang {{ activeStatus }}</h3>
        <table class="summary-table">
          <thead>
            <tr>
              <th>No</th>
              <th>Nama</th>
              <th>Merk</th>
              <th>Status</th>
              <th>Deskripsi</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(item, index) in filteredBarang"
              :key="index"
            >
              <td>{{ index + 1 }}</td>
              <td>{{ item.nama }}</td>
              <td>{{ item.merk }}</td>
              <td>{{ item.status }}</td>
              <td>{{ item.deskripsi }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'DaftarBarangPage',
  data() {
    return {
      activeStatus: '',
      barangList: [
        { nama: 'Laptop', merk: 'Asus', status: 'Dipinjam', deskripsi: 'Digunakan presentasi' },
        { nama: 'Proyektor', merk: 'Epson', status: 'Tersedia', deskripsi: 'Baru' },
        { nama: 'Kursi Guru', merk: '-', status: 'Rusak', deskripsi: 'Kaki patah' },
        { nama: 'Monitor', merk: 'LG', status: 'Tersedia', deskripsi: '-' },
        { nama: 'Speaker', merk: 'Polytron', status: 'Rusak', deskripsi: 'Tidak berbunyi' }
      ]
    };
  },
  computed: {
    filteredBarang() {
      return this.barangList.filter(b => b.status === this.activeStatus);
    }
  },
  methods: {
    setStatus(status) {
      this.activeStatus = status;
    }
  }
};
</script>

<style scoped>
.main-content {
  flex: 1;
  background-color: #f4f4f4;
  display: flex;
  flex-direction: column;
}

.header {
  background-color: white;
  padding: 20px;
  border-bottom: 1px solid #ddd;
}

.content {
  padding: 20px;
}

.cards {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  margin-top: 20px;
}

.card {
  background-color: white;
  padding: 20px;
  flex: 1 1 250px;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  text-align: center;
  cursor: pointer;
  transition: background-color 0.2s;
}

.card:hover {
  background-color: #ecf0f1;
}

.card h3 {
  margin-bottom: 10px;
  color: #2c3e50;
}

.card p {
  font-size: 24px;
  font-weight: bold;
}

.card.available p {
  color: #27ae60;
}
.card.damaged p {
  color: #e74c3c;
}
.card.borrowed p {
  color: #f39c12;
}

.table-section {
  margin-top: 30px;
}

.close-btn {
  background-color: #e74c3c;
  color: white;
  border: none;
  padding: 6px 12px;
  border-radius: 4px;
  margin-bottom: 10px;
  cursor: pointer;
}

.summary-table {
  width: 100%;
  border-collapse: collapse;
  background-color: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

.summary-table th,
.summary-table td {
  padding: 12px 16px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

.summary-table thead {
  background-color: #2c3e50;
  color: white;
}

.summary-table tbody tr:hover {
  background-color: #f1f1f1;
}
</style>
