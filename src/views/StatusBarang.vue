<template>
  <div class="status-barang">
    <h2>Status Barang</h2>
    <div class="status-cards">
      <div
        v-for="status in statusList"
        :key="status.nama"
        class="status-card"
        @click="filterStatus(status.nama)"
      >
        <h3>{{ status.nama }}</h3>
        <p>{{ status.jumlah }} Barang</p>
      </div>
    </div>

    <div class="table-container" v-if="filteredBarang.length > 0">
      <h3>Daftar Barang: {{ selectedStatus }}</h3>
      <table class="status-table">
        <thead>
          <tr>
            <th>No</th>
            <th>Nama Barang</th>
            <th>Merk</th>
            <th>Status</th>
            <th>Deskripsi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(barang, index) in filteredBarang" :key="barang.id">
            <td>{{ index + 1 }}</td>
            <td>{{ barang.nama }}</td>
            <td>{{ barang.merk }}</td>
            <td>{{ barang.status }}</td>
            <td>{{ barang.deskripsi }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: 'StatusBarang',
  data() {
    return {
      statusList: [
        { nama: 'Tersedia', jumlah: 30 },
        { nama: 'Dipinjam', jumlah: 10 },
        { nama: 'Rusak', jumlah: 5 },
        { nama: 'Service', jumlah: 2 },
        { nama: 'Hilang', jumlah: 1 },
      ],
      allBarang: [
        { id: 1, nama: 'Laptop', merk: 'Asus', status: 'Dipinjam', deskripsi: 'Digunakan untuk presentasi' },
        { id: 2, nama: 'Proyektor', merk: 'Epson', status: 'Tersedia', deskripsi: '-' },
        { id: 3, nama: 'Meja Guru', merk: '-', status: 'Rusak', deskripsi: 'Kaki patah' },
        // dst.
      ],
      selectedStatus: '',
    };
  },
  computed: {
    filteredBarang() {
      if (!this.selectedStatus) return [];
      return this.allBarang.filter(barang => barang.status === this.selectedStatus);
    }
  },
  methods: {
    filterStatus(status) {
      this.selectedStatus = status;
    }
  }
};
</script>

<style scoped>
.status-cards {
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
  margin: 20px 0;
}

.status-card {
  background-color: #f8f9fa;
  border-left: 5px solid #2980b9;
  padding: 20px;
  border-radius: 6px;
  cursor: pointer;
  flex: 1 1 200px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  transition: 0.2s;
}

.status-card:hover {
  background-color: #ecf0f1;
}

.table-container {
  overflow-x: auto;
  background-color: white;
  margin-top: 20px;
  border-radius: 6px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.05);
}

.status-table {
  width: 100%;
  min-width: 800px;
  border-collapse: collapse;
}

.status-table th, .status-table td {
  padding: 12px;
  border-bottom: 1px solid #ccc;
}

.status-table th {
  background-color: #2c3e50;
  color: white;
}
</style>
