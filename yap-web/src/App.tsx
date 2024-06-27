import React, { useState } from 'react';
import './App.css';
import { authClient } from './clients';
import HomePage from "./Homepage";

function App() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [currentRoom, setCurrentRoom] = useState('');

  const handleLogin = async () => {
    setLoading(true);
    setError('');
    try {
      const response = await authClient.login({ username, password });
      console.log('Login successful:', response);
      setIsLoggedIn(true);
    } catch (error) {
      setError('Login failed. Please try again.');
      console.error('Login failed:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleCreateAccount = async () => {
    setLoading(true);
    setError('');
    try {
      const response = await authClient.register({ username, password });
      console.log('Account creation successful:', response);
      setIsLoggedIn(true);
    } catch (error) {
      setError('Account creation failed. Please try again.');
      console.error('Account creation failed:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    setUsername('');
    setPassword('');
    setCurrentRoom('');
  };

  const handleJoinChat = (roomName: string) => {
    setCurrentRoom(roomName);
    // Here you would typically set up the chat room connection
    console.log(`Joined chat room: ${roomName}`);
  };

  if (isLoggedIn) {
    return (
      <HomePage
        username={username}
        onLogout={handleLogout}
        onJoinChat={handleJoinChat}
      />
    );
  }

  return (
    <div className="App">
      <div className="welcome-panel">
        <h2>Welcome!</h2>
        <p>Please sign in to your account or create a new one.</p>
      </div>
      <div className="login-form">
        <input
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          placeholder="Username"
        />
        <input
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="Password"
        />
        {error && <div className="error">{error}</div>}
        <div className="button-row">
          <button onClick={handleLogin} disabled={loading}>Login</button>
          <button onClick={handleCreateAccount} disabled={loading}>Create Account</button>
        </div>
        {loading && <div className="loading">Loading...</div>}
      </div>
    </div>
  );
}

export default App;