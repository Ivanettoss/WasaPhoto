<template>

<header class="header">
    <a class="logo">WasaPhoto</a>
    <nav class="navbar">
      <div class="search-container">

        <input id="searchForm" type="text"
        placeholder="Search users..."
        v-model="searchQuery" 
        @keyup.enter="searchUsers(searchQuery)"  
      />	
      <button class="search-button" @click="searchUsers(searchQuery)"></button>

      <div v-if="users.length > 0" class="search-dropdown">
        <ul>
          <li v-for="user in users" :key="user.username" @click="goToProfile(user.username)">
            {{ user.username }} 
          </li>
        </ul>
      </div>
  </div>


        <a  id="menubu">Home</a>
        <a  id="menubu">Profile</a>
        <a  id="menubu" @click="settingsDropdown">Settings</a>

        <div class="dropdown">
         <div v-if="showSettings" class="dropdown-content">

        <a  id="menubu" @click="showPopup=true">Change Username</a>
        <a   id="menubu" @click="doLogOut()">Logout</a>
      </div>
      </div>
    </nav>
</header>
