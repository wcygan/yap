import React, { useState } from 'react';
import logo from './logo.svg';
import './App.css';

import client from './connectClient';

function App() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    // Call gRPC service for login
  };

  const handleCreateAccount = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    // Call gRPC service for account creation
  };

  return (
    <div>
      <div className="login-form">
        <form onSubmit={handleLogin}>
          <input type="text" value={username} onChange={(e) => setUsername(e.target.value)} placeholder="Username" />
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} placeholder="Password" />
          <button type="submit">Login</button>
        </form>
        <form onSubmit={handleCreateAccount}>
          <input type="text" value={username} onChange={(e) => setUsername(e.target.value)} placeholder="Username" />
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} placeholder="Password" />
          <button type="submit">Create Account</button>
        </form>
      </div>
    </div>
  );
}

export default App;