import { BrowserRouter, Route, Routes } from "react-router-dom";
import Sidebar from "./component/sidebar"; // Make sure this path is correct
import Dashboard from "./pages/dashboard/dashboard"; // Make sure this path is correct
import BookingManagement from "./pages/booking/managebook";
import RoomManagement from "./pages/room/roommangement";

function App() {
  return (
    <BrowserRouter>
      <Sidebar /> {/* Sidebar renders here */}
      <div style={{ marginLeft: 240 }}> {/* Space for the sidebar */}
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/bookmanage" element={<BookingManagement />} />
          <Route path="/roommanage" element={<RoomManagement />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
