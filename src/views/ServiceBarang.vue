<template>
  <div class="service-barang">
    <h2>Service Barang</h2>

    <button class="btn-show-form" @click="showForm = !showForm">
      {{ showForm ? 'Tutup Form' : '+ Tambah Barang ke Service' }}
    </button>

    <div v-if="showForm" class="form-service">
      <h3>Tambah Barang ke Service</h3>
      <form @submit.prevent="tambahService">
        <div class="form-row">
          <label>Nama Barang:</label>
          <select v-model="form.barang_id" required>
            <option disabled value="">Pilih Barang</option>
            <option v-for="b in barangList" :key="b.id" :value="b.id">
              {{ b.nama }}
            </option>
          </select>
        </div>
        <div class="form-row">
          <label>Keterangan Kerusakan:</label>
          <input v-model="form.kerusakan" type="text" required />
        </div>
        <div class="form-row">
          <label>Tanggal Masuk:</label>
          <input v-model="form.tanggalMasuk" type="date" required />
        </div>
        <button type="submit" class="btn-tambah">Simpan</button>
      </form>
    </div>

    <div class="table-container">
      <h3>Riwayat Service Barang</h3>
      <table class="service-table">
        <thead>
          <tr>
            <th>No</th>
            <th>Nama Barang</th>
            <th>Kerusakan</th>
            <th>Tanggal Masuk</th>
            <th>Status</th>
            <th>Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in serviceList" :key="item.id">
            <td>{{ index + 1 }}</td>
            <td>{{ item.nama_barang }}</td>
            <td>{{ item.deskripsi }}</td>
            <td>{{ item.tanggal_service }}</td>
            <td>{{ item.status }}</td>
            <td>
              <button @click="selesaiService(item.id)" class="btn-selesai">Selesai</button>
              <button @click="hapusService(item.id)" class="btn-delete">Hapus</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: "ServiceBarang",
  data() {
    return {
      showForm: false,
      form: {
        barang_id: '',
        kerusakan: '',
        tanggalMasuk: '',
      },
      serviceList: [],
      barangList: []
    }
  },
  mounted() {
    this.getServiceList();
    this.ambilBarangList();
  },
  methods: {
    getServiceList() {
      fetch("http://localhost:3000/service")
        .then(res => res.json())
        .then(data => {
          this.serviceList = data;
        })
        .catch(() => alert("Gagal mengambil data service."));
    },
    async ambilBarangList() {
  try {
    const res = await fetch("http://localhost:3000/barang");
    const data = await res.json();
    this.barangList = data.filter(b => b.status === "Rusak");
  } catch (err) {
    console.error("Gagal mengambil data barang:", err);
  }
},
    tambahService() {
      const payload = {
        barang_id: this.form.barang_id,
        deskripsi: this.form.kerusakan,
        tanggal_service: this.form.tanggalMasuk,
        status: "Dalam Service"
      };
      fetch("http://localhost:3000/service", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
      })
        .then(() => {
          this.getServiceList();
          this.form = { barang_id: '', kerusakan: '', tanggalMasuk: '' };
          this.showForm = false;
        })
        .catch(() => alert("Gagal menambah data."));
    },
    selesaiService(id) {
  fetch(`http://localhost:3000/service/selesai?id=${id}`, {
    method: "PUT"
  })
    .then(() => this.getServiceList())
    .catch(() => alert("Gagal update status."));
},
    hapusService(id) {
      if (!confirm("Hapus data service ini?")) return;
      fetch(`http://localhost:3000/service?id=${id}`, {
        method: "DELETE"
      })
        .then(() => this.getServiceList())
        .catch(() => alert("Gagal menghapus data."));
    }
  }
}
</script>

<style scoped>
.service-barang {
  padding: 20px;
}
.btn-show-form {
  background-color: #2c3e50;
  color: white;
  padding: 10px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-bottom: 15px;
}
.btn-show-form:hover {
  background-color: #1a252f;
}
.form-service {
  background-color: white;
  padding: 20px;
  margin-bottom: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}
.form-row {
  margin-bottom: 12px;
  display: flex;
  flex-direction: column;
}
label {
  margin-bottom: 4px;
  font-weight: bold;
}
input {
  padding: 8px;
  border-radius: 4px;
  border: 1px solid #ccc;
}
.btn-tambah {
  background-color: #27ae60;
  color: white;
  padding: 10px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-top: 10px;
}
.btn-tambah:hover {
  background-color: #219150;
}
.table-container {
  overflow-x: auto;
  background: white;
  padding: 10px;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.05);
}
.service-table {
  width: 100%;
  min-width: 900px;
  border-collapse: collapse;
}
.service-table th,
.service-table td {
  padding: 12px;
  border-bottom: 1px solid #ddd;
}
.service-table th {
  background-color: #2c3e50;
  color: white;
}
.btn-selesai {
  background-color: #3498db;
  color: white;
  padding: 6px 10px;
  border-radius: 4px;
  border: none;
  margin-right: 5px;
}
.btn-delete {
  background-color: #e74c3c;
  color: white;
  padding: 6px 10px;
  border-radius: 4px;
  border: none;
}
.btn-selesai:hover {
  background-color: #2980b9;
}
.btn-delete:hover {
  background-color: #c0392b;
}
</style>
