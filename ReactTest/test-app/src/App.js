import Navbar from "./Components/Navbar";
import HomePage from "./Components/HomePage";
import { useState } from 'react';


function App() {
  const [JwtToken, SetJwtToken] = useState('');
  const [IsLoggedIn, SetIsLoggedIn] = useState(false);
  const [Route, SetRoute] = useState('/')

  return (
    <div className="App">
      <Navbar IsLoggedIn={IsLoggedIn}/>
      {Route === '/' && <HomePage/>}
    </div>
  );
}

export default App;
