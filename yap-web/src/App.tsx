import React, { useState } from 'react';
import './App.css';
import { authClient } from './clients';

function App() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = async () => {
    try {
      const response = await authClient.login({ username, password });
      console.log('Login successful:', response);
    } catch (error) {
      console.error('Login failed:', error);
    }
  };

  const handleCreateAccount = async () => {
    try {
      const response = await authClient.register({ username, password });
      console.log('Account creation successful:', response);
    } catch (error) {
      console.error('Account creation failed:', error);
    }
  };

  return (
    <div>
      <div className="login-form">
        <input type="text" value={username} onChange={(e) => setUsername(e.target.value)} placeholder="Username" />
        <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} placeholder="Password" />
        <div className="button-row">
          <button onClick={handleLogin}>Login</button>
          <button onClick={handleCreateAccount}>Create Account</button>
        </div>
      </div>
    </div>
  );
}

export default App;