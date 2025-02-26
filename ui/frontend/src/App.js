import React, { useState } from 'react';
import './App.css';
import axios from 'axios';

function App() {
  const [participants, setParticipants] = useState(['']);
  const [outcomes, setOutcomes] = useState(['']);
  const [results, setResults] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleParticipantChange = (index, value) => {
    const newParticipants = [...participants];
    newParticipants[index] = value;
    setParticipants(newParticipants);
  };

  const handleOutcomeChange = (index, value) => {
    const newOutcomes = [...outcomes];
    newOutcomes[index] = value;
    setOutcomes(newOutcomes);
  };

  const addParticipant = () => {
    setParticipants([...participants, '']);
    // 참가자가 추가될 때 결과도 함께 추가
    if (outcomes.length < participants.length + 1) {
      setOutcomes([...outcomes, '']);
    }
  };

  const addOutcome = () => {
    setOutcomes([...outcomes, '']);
    // 결과가 추가될 때 참가자도 함께 추가
    if (participants.length < outcomes.length + 1) {
      setParticipants([...participants, '']);
    }
  };

  const removeParticipant = (index) => {
    if (participants.length > 1) {
      const newParticipants = [...participants];
      newParticipants.splice(index, 1);
      setParticipants(newParticipants);
    }
  };

  const removeOutcome = (index) => {
    if (outcomes.length > 1) {
      const newOutcomes = [...outcomes];
      newOutcomes.splice(index, 1);
      setOutcomes(newOutcomes);
    }
  };

  const runLadderGame = async () => {
    // 입력 검증
    const filteredParticipants = participants.filter(p => p.trim() !== '');
    const filteredOutcomes = outcomes.filter(o => o.trim() !== '');

    if (filteredParticipants.length < 2) {
      setError('최소 2명의 참가자가 필요합니다.');
      return;
    }

    if (filteredOutcomes.length < 2) {
      setError('최소 2개의 결과가 필요합니다.');
      return;
    }

    if (filteredParticipants.length !== filteredOutcomes.length) {
      setError('참가자 수와 결과 수가 일치해야 합니다.');
      return;
    }

    // 중복 검사
    const uniqueParticipants = new Set(filteredParticipants);
    if (uniqueParticipants.size !== filteredParticipants.length) {
      setError('참가자 이름은 중복될 수 없습니다.');
      return;
    }

    setError('');
    setLoading(true);

    try {
      const response = await axios.post('/ladder', {
        participants: filteredParticipants,
        outcomes: filteredOutcomes
      });
      
      setResults(response.data.results);
    } catch (err) {
      setError('사다리 게임 실행 중 오류가 발생했습니다: ' + (err.response?.data?.error || err.message));
    } finally {
      setLoading(false);
    }
  };

  const resetGame = () => {
    setParticipants(['']);
    setOutcomes(['']);
    setResults(null);
    setError('');
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>사다리 게임</h1>
      </header>
      <main className="App-main">
        {!results ? (
          <div className="setup-container">
            <div className="input-section">
              <h2>참가자</h2>
              {participants.map((participant, index) => (
                <div key={`participant-${index}`} className="input-group">
                  <input
                    type="text"
                    value={participant}
                    onChange={(e) => handleParticipantChange(index, e.target.value)}
                    placeholder="참가자 이름"
                  />
                  <button 
                    className="remove-btn"
                    onClick={() => removeParticipant(index)}
                    disabled={participants.length <= 1}
                  >
                    -
                  </button>
                </div>
              ))}
              <button className="add-btn" onClick={addParticipant}>참가자 추가</button>
            </div>

            <div className="input-section">
              <h2>결과</h2>
              {outcomes.map((outcome, index) => (
                <div key={`outcome-${index}`} className="input-group">
                  <input
                    type="text"
                    value={outcome}
                    onChange={(e) => handleOutcomeChange(index, e.target.value)}
                    placeholder="결과"
                  />
                  <button 
                    className="remove-btn"
                    onClick={() => removeOutcome(index)}
                    disabled={outcomes.length <= 1}
                  >
                    -
                  </button>
                </div>
              ))}
              <button className="add-btn" onClick={addOutcome}>결과 추가</button>
            </div>

            {error && <div className="error-message">{error}</div>}

            <button 
              className="run-btn" 
              onClick={runLadderGame}
              disabled={loading}
            >
              {loading ? '처리 중...' : '사다리 게임 실행'}
            </button>
          </div>
        ) : (
          <div className="results-container">
            <h2>게임 결과</h2>
            <div className="results-list">
              {Object.entries(results).map(([participant, outcome], index) => (
                <div key={index} className="result-item">
                  <span className="participant">{participant}</span>
                  <span className="arrow">→</span>
                  <span className="outcome">{outcome}</span>
                </div>
              ))}
            </div>
            <button className="reset-btn" onClick={resetGame}>다시 시작</button>
          </div>
        )}
      </main>
    </div>
  );
}

export default App;
