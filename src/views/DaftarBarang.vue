<template>
    <div class="main-content">
        <header class="header">
            <h1>Daftar Barang</h1>
        </header>

        <section class="content">
            <h2>Ringkasan Status Barang</h2>
            <div class="cards">
                <div class="card available">
                    <h3>Barang Tersedia</h3>
                    <p>80</p>
                </div>
                <div class="card damaged">
                    <h3>Barang Rusak</h3>
                    <p>25</p>
                </div>
                <div class="card borrowed">
                    <h3>Barang Dipinjam</h3>
                    <p>15</p>
                </div>
            </div>

            <!-- tempatkan konten tambahan disini seperti tabel -->
        </section>
    </div>
</template>

<script>
export default {
    name: 'DaftarBarangPage',
}
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
}

.card h3 {
    margin-bottom: 10px;
    color: #2c3e50;
}

.card p {
    font-size: 24px;
    font-weight: bold;
}

/* warna khusus per status */
.card.available p {
    color: #27ae60; 
}

.card.damage p {
    color: #e74c3c;
}

.card.borrowed p {
    color: #f39c12;
}
</style>