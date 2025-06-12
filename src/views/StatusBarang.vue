<template>
  <div class="status-barang">
    <h2>Status Barang</h2>

    <div class="status-cards">
      <div
        v-for="status in statusList || []"
        :key="status.status"
        class="status-card"
        @click="filterStatus(status.status)"
      >
        <h3>{{ status.status }}</h3>
        <p>{{ status.jumlah }}Barang</p>
      </div>
    </div>

    <div class="table-container" v-if="selectedStatus && filteredBarang.length" >
      <button class="close-btn" @click="selectedStatus = ''">Tutup</button>
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
      statusList: [],
      allBarang: [],
      selectedStatus: ''
    };
  },
  computed: {
    filteredBarang() {
      return this.allBarang;
    }
  },
  mounted() {
    this.fetchStatusList();
  },
  methods: {
    fetchStatusList() {
      fetch('http://localhost:3000/status')
        .then(res => res.json())
        .then(data => {
          this.statusList = data;
        })
        .catch(err => console.error("Gagal ambil data status:", err));
    },
    filterStatus(status) {
      this.selectedStatus = status;
      fetch(`http://localhost:3000/status/filter?status=${status}`)
        .then(res => res.json())
        .then(data => {
          this.allBarang = data;
        })
        .catch(err => console.error("Gagal ambil data barang:", err));
    }
  }
};
</script>

<style scoped>
.status-barang {
  padding: 16px;
}

.status-cards {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
  justify-content: center;
  margin: 20px 0;
}

.status-card {
  background-color: #f8f9fa;
  border-left: 5px solid #2980b9;
  padding: 16px;
  border-radius: 8px;
  cursor: pointer;
  flex: 1 1 180px;
  max-width: 250px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  transition: background-color 0.3s ease;
}

.status-card:hover {
  background-color: #ecf0f1;
}

.table-container {
  overflow-x: auto;
  margin-top: 20px;
  border-radius: 8px;
  background-color: white;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

.status-table {
  width: 100%;
  min-width: 700px;
  border-collapse: collapse;
}

.status-table th, .status-table td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

.status-table th {
  background-color: #2c3e50;
  color: white;
}

.close-btn {
  background-color: #e74c3c;
  color: white;
  border: none;
  padding: 6px 12px;
  border-radius: 4px;
  cursor: pointer;
  margin-bottom: 10px;
}

.close-btn:hover {
  background-color: #c0392b;
}

@media (max-width: 600px) {
  .status-cards {
    flex-direction: column;
    align-items: stretch;
  }

  .status-card {
    flex: 1 1 auto;
    max-width: 100%;
  }

  .status-table {
    font-size: 14px;
  }

  .table-container {
    display: block;
  }
}
</style>
