import React, { useState, useEffect } from "react";
import { Table, Tag, Button } from "antd";
import { Dialog, DialogTitle, DialogContent, DialogActions, Button as MUIButton } from "@mui/material"; // MUI Button ถูกปรับมาใช้จาก MUIButton
import "./RoomManagement.css";

// กำหนด type ให้กับห้องพัก
interface Room {
  id: number;
  roomNumber: string;
  roomType: string;
  price: number;
  status: string;
}

const RoomManagement = () => {
  const [selectedRoom, setSelectedRoom] = useState<Room | null>(null);
  const [openDialog, setOpenDialog] = useState(false);
  const [rooms, setRooms] = useState<Room[]>([]);

  // โหลดข้อมูลเมื่อ component ถูก mount
  useEffect(() => {
    setTimeout(() => {
      setRooms([
        { id: 1, roomNumber: "101", roomType: "Deluxe", price: 1500, status: "Available" },
        { id: 2, roomNumber: "102", roomType: "Suite", price: 2500, status: "Booked" },
        { id: 3, roomNumber: "103", roomType: "Standard", price: 1000, status: "Available" },
      ]);
    }, 500); // จำลองการโหลดข้อมูล
  }, []);

  // คอลัมน์ของตาราง
  const columns = [
    { title: "หมายเลขห้อง", dataIndex: "roomNumber", key: "roomNumber" },
    { title: "ประเภท", dataIndex: "roomType", key: "roomType" },
    { title: "ราคา (฿)", dataIndex: "price", key: "price", render: (price: number) => `฿${price.toLocaleString()}` },
    {
      title: "สถานะ",
      dataIndex: "status",
      key: "status",
      render: (status: string) => (
        <Tag color={status === "Available" ? "green" : "red"}>{status}</Tag>
      ),
    },
    {
      title: "จัดการ",
      key: "action",
      render: (_: any, record: Room) => (
        <div className="action-buttons">
          <Button onClick={() => handleEdit(record)} type="primary">แก้ไข</Button>
          <Button danger onClick={() => handleDelete(record.id)}>ลบ</Button>
        </div>
      ),
    },
  ];

  // ฟังก์ชันสำหรับแก้ไขข้อมูลห้อง
  const handleEdit = (room: Room) => {
    setSelectedRoom(room);
    setOpenDialog(true);
  };

  // ฟังก์ชันสำหรับลบห้อง
  const handleDelete = (roomId: number) => {
    setRooms(rooms.filter(room => room.id !== roomId));
  };

  // ฟังก์ชันสำหรับปิด Dialog
  const handleCloseDialog = () => {
    setOpenDialog(false);
    setSelectedRoom(null); // รีเซ็ตข้อมูลห้องเมื่อปิด Dialog
  };

  return (
    <div className="room-management-container">
      <h2 className="room-management-title">จัดการห้องพัก</h2>
      
      {/* ถ้าข้อมูลห้องว่างเปล่า จะมีข้อความแจ้ง */}
      {rooms.length === 0 ? (
        <p>กำลังโหลดข้อมูลห้องพัก...</p>
      ) : (
        <Table
          dataSource={rooms}
          columns={columns}
          rowKey="id"
          pagination={{ pageSize: 5 }}
          className="room-table"
        />
      )}

      {/* Dialog สำหรับแก้ไขห้องพัก */}
      <Dialog 
        open={openDialog} 
        onClose={handleCloseDialog} 
        maxWidth="lg" // ขยาย Dialog ให้กว้างขึ้น
        sx={{
          "& .MuiDialog-paper": {
            padding: "30px", // เพิ่ม padding ภายใน Dialog
            width: "80%", // ให้ Dialog ขยายเต็ม 80% ของหน้าจอ
            maxWidth: "650px", // ขนาดสูงสุดที่ 650px
            maxHeight: "80vh", // ขนาดสูงสุดของ Dialog
            borderRadius: "12px", // ขอบมุมมน
            boxShadow: "0 4px 16px rgba(0, 0, 0, 0.2)", // เพิ่มเงามุมที่นุ่มนวล
          }
        }}
      >
        <DialogTitle sx={{ fontSize: "26px", fontWeight: "bold", textAlign: "center" }}>แก้ไขข้อมูลห้องพัก</DialogTitle>
        <DialogContent sx={{ fontSize: "18px" }}>
          {selectedRoom ? (
            <>
              <p><strong>หมายเลขห้อง:</strong> {selectedRoom.roomNumber}</p>
              <p><strong>ประเภท:</strong> {selectedRoom.roomType}</p>
              <p><strong>ราคา:</strong> ฿{selectedRoom.price.toLocaleString()}</p>
            </>
          ) : (
            <p>ไม่พบข้อมูล</p>
          )}
        </DialogContent>
        <DialogActions sx={{ justifyContent: "center" }}>
          <MUIButton onClick={handleCloseDialog} color="primary" sx={{ fontSize: "16px", padding: "10px 20px", margin: "0 10px" }}>ปิด</MUIButton>
          <MUIButton onClick={handleCloseDialog} color="success" sx={{ fontSize: "16px", padding: "10px 20px", margin: "0 10px" }}>บันทึก</MUIButton>
        </DialogActions>
      </Dialog>
    </div>
  );
};

export default RoomManagement;
