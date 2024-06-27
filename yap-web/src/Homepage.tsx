import React, {useState} from 'react';
import {chatRoomClient} from './clients';

interface HomePageProps {
    username: string;
    onLogout: () => void;
    onJoinChat: (roomName: string) => void;
}

const HomePage: React.FC<HomePageProps> = ({username, onLogout, onJoinChat}) => {
    const [channelName, setChannelName] = useState('');
    const [error, setError] = useState('');
    const [loading, setLoading] = useState(false);

    const handleJoinRoom = async () => {
        if (!channelName) {
            setError('Please enter a room name.');
            return;
        }
        setLoading(true);
        setError('');
        try {
            chatRoomClient.joinChatRoom({userId: username, userName: username, channelName: channelName});
            onJoinChat(channelName);
        } catch (error) {
            setError('Failed to join chat room. Please try again.');
            console.error('Chat room join failed:', error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="App">
            <div className="welcome-panel">
                <h2>Welcome, {username}!</h2>
                <p>Create a new chat room or join an existing one.</p>
            </div>
            <div className="login-form">
                <input
                    type="text"
                    value={channelName}
                    onChange={(e) => setChannelName(e.target.value)}
                    placeholder="Chat Room Name"
                />
                {error && <div className="error">{error}</div>}
                <div className="button-row">
                    <button onClick={handleJoinRoom} disabled={loading}>Join Room</button>
                </div>
                <button onClick={onLogout} disabled={loading}>Logout</button>
                {loading && <div className="loading">Loading...</div>}
            </div>
        </div>
    );
};
export default HomePage;