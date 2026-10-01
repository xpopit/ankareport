<template>
  <div class="app-layout">
    <nav class="sidebar">
      <h2>XERP Report Studio</h2>
      <ul>
        <li><router-link to="/">Dashboard</router-link></li>
        <li><router-link to="/reports">Reports</router-link></li>
        <li><router-link to="/datasets">Datasets</router-link></li>
        <li><router-link to="/data-sources">Data Sources</router-link></li>
      </ul>
    </nav>
    <main class="content">
      <router-view></router-view>
    </main>
  </div>
</template>

<style scoped>
.app-layout {
  display: flex;
  height: 100vh;
}
.sidebar {
  width: 250px;
  background: #2c3e50;
  color: white;
  padding: 20px;
}
.sidebar h2 {
  margin-top: 0;
}
.sidebar ul {
  list-style: none;
  padding: 0;
}
.sidebar li {
  margin-bottom: 10px;
}
.sidebar a {
  color: white;
  text-decoration: none;
}
.content {
  flex: 1;
  padding: 20px;
  overflow: auto;
}
</style>
