<template>
    <div class="login-page">
        <div class="login-card card shadow-sm p-4">
            <h4 class="text-center mb-3">Login admin</h4>
            <form @submit.prevent="handleLogin">
                <div class="mb-3">
                    <label for="username" class="form-label">username</label>
                    <input v-model="username" type="username" id="username" class="form-control" required />
                </div>
                <div class="mb-3">
                    <label for="password" class="form-label">Password</label>
                    <input v-model="password" type="password" id="password" class="form-control" required />
                </div>
                <button type="submit" class="btn btn-primary w-100">Login</button>
                <p v-if="error" class="text-danger mt-2">{{ error }}</p>
            </form>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'LoginPage',
  data() {
    return {
      username: '', 
      password: '',
      error: ''
    }
  },
  methods: {
    async handleLogin() {
      try {
        const response = await axios.post('http://localhost:3000/auth/login', {
          username: this.username,
          password: this.password
        })

        const { user, token } = response.data

        // Simpan token dan user info (opsional)
        localStorage.setItem('token', token)
        localStorage.setItem('user', JSON.stringify(user))

        // Arahkan ke halaman user atau dashboard
        this.$router.push('/dashboard')
      } catch (error) {
        this.error = 'Username atau password salah.'
        console.error('Login error:', error)
      }
    }
  }
}
</script>


<style scoped>
.login-page {
    min-height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: #f8f9fa;
    padding: 1rem;
}

.login-card {
    width: 100%;
    max-width: 320px;
    background-color: white;
    border-radius: 8px;
}

</style>