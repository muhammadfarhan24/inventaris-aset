<template>
  <div class="main-content">
    <header class="header">
      <h1>Daftar Barang</h1>
    </header>

    <section class="content">
      <!-- Form Tambah Barang -->
      <h2>Tambah Barang</h2>
      <form @submit.prevent="tambahBarang" class="form-barang">
        <input v-model="form.nama" placeholder="Nama Barang" required />
        <input v-model="form.kategori_id" placeholder="Kategori ID" required />
        <input v-model="form.merk_id" placeholder="Merk ID" required />
        <select v-model="form.status" required>
          <option disabled value="">Pilih Status</option>
          <option value="Tersedia">Tersedia</option>
          <option value="Rusak">Rusak</option>
          <option value="Dipinjam">Dipinjam</option>
        </select>
        <select v-model="form.ruangan_id" required>
          <option disabled value="">Pilih Ruangan</option>
          <option v-for="ruangan in ruanganList" :key="ruangan.id" :value="ruangan.id">
            {{ ruangan.nama }}
          </option>
        </select>
        <input v-model="form.deskripsi" placeholder="Deskripsi" />
        <button type="submit">Tambah</button>
      </form>

      <!-- Ringkasan Status -->
      <h2>Ringkasan Status Barang</h2>
      <div class="cards">
        <div class="card available" @click="setStatus('Tersedia')">
          <h3>Barang Tersedia</h3>
          <p>{{ jumlahStatus('Tersedia') }}</p>
        </div>
        <div class="card damaged" @click="setStatus('Rusak')">
          <h3>Barang Rusak</h3>
          <p>{{ jumlahStatus('Rusak') }}</p>
        </div>
        <div class="card borrowed" @click="setStatus('Dipinjam')">
          <h3>Barang Dipinjam</h3>
          <p>{{ jumlahStatus('Dipinjam') }}</p>
        </div>
      </div>

      <!-- Tabel berdasarkan status -->
      <div v-if="activeStatus" class="table-section">
        <button class="close-btn" @click="activeStatus = ''">Tutup</button>
        <h3>Daftar Barang {{ activeStatus }}</h3>
        <table class="summary-table">
          <thead>
            <tr>
              <th>No</th>
              <th>Nama</th>
              <th>Kategori</th>
              <th>Merk</th>
              <th>Status</th>
              <th>Ruangan</th>
              <th>Deskripsi</th>
              <th>Creted_at</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in filteredBarang" :key="index">
              <td>{{ index + 1 }}</td>
              <td>{{ item.nama }}</td>
              <td>{{ item.kategori_id }}</td>
              <td>{{ item.merk_id }}</td>
              <td>{{ item.status }}</td>
              <td>{{ item.ruangan_id }}</td>
              <td>{{ item.deskripsi }}</td>
              <td>{{ item.created_at }}</td>
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
      barangList: [],
      ruanganList:[],
      form: {
        nama: '',
        kategori_id: '',
        merk_id: '',
        status: '',
        ruangan_id: null,
        deskripsi: ''
      }
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
    },
    jumlahStatus(status) {
      return this.barangList.filter(b => b.status === status).length;
    },
    async ambilDataBarang() {
      try {
        const res = await fetch('http://localhost:3000/barang');
        const data = await res.json();
        this.barangList = data;
      } catch (err) {
        console.error('Gagal ambil data barang:', err);
      }
    },
    async ambilDataRuangan() {
  try {
    const res = await fetch('http://localhost:3000/ruangan');
    const data = await res.json();
    this.ruanganList = data;
  } catch (err) {
    console.error('Gagal ambil data ruangan:', err);
  }
},
    async tambahBarang() {
      try {
        const res = await fetch('http://localhost:3000/barang', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(this.form)
        });
        if (res.ok) {
          this.form = {
            nama: '',
            kategori_id: '',
            merk_id: '',
            status: '',
            ruangan_id: '',
            deskripsi: ''
          };
          this.ambilDataBarang(); // refresh list
        } else {
          console.error('Gagal tambah barang');
        }
      } catch (err) {
        console.error('Gagal koneksi ke server:', err);
      }
    }
  },
  mounted() {
    this.ambilDataBarang();
    this.activeStatus = 'Tersedia';
    this.ambilDataRuangan();
  }
};
</script>

<style scoped>
.main-content {
  background-color: #f4f4f4;
  padding: 20px;
}
.header {
  background-color: white;
  padding: 20px;
  border-bottom: 1px solid #ddd;
}
.cards {
  display: flex;
  gap: 20px;
  margin-top: 20px;
}
.card {
  flex: 1;
  background: white;
  padding: 20px;
  border-radius: 8px;
  cursor: pointer;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}
.card.available p { color: #27ae60; }
.card.damaged p { color: #e74c3c; }
.card.borrowed p { color: #f39c12; }
.table-section {
  margin-top: 30px;
}
.summary-table {
  width: 100%;
  border-collapse: collapse;
  background-color: white;
  margin-top: 10px;
}
.summary-table th, .summary-table td {
  padding: 10px;
  border-bottom: 1px solid #ddd;
}
.summary-table thead {
  background-color: #2c3e50;
  color: white;
}
.close-btn {
  background-color: #e74c3c;
  color: white;
  padding: 6px 12px;
  border: none;
  margin-bottom: 10px;
  border-radius: 4px;
}
.form-barang {
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-bottom: 30px;
}
.form-barang input, .form-barang select {
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
}
.form-barang button {
  width: fit-content;
  padding: 8px 16px;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
</style>
