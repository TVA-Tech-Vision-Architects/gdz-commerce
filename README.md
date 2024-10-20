# gdz-commerce
development
feature-feature
* feature/feature
main
stagingTest

แนวทางการทำงานกับ(branch) ของ Git ในทีมพัฒนาเป็นสิ่งสำคัญในการจัดการเวิร์กโฟลว์ให้เป็นระเบียบดังนั้น แต่ละ(branch)ควรมีหน้าที่และความรับผิดชอบที่ชัดเจนเพื่อให้ทีมทำงานได้
1. main (สาขาหลัก)
สาขา main ควรเป็นสาขาที่มีโค้ดเวอร์ชันเสถียรที่สุดและสามารถใช้งานได้จริงในโปรดักชัน (Production)
การเปลี่ยนแปลงที่รวม (merge) เข้ามาใน main จะต้องผ่านการทดสอบและตรวจสอบเรียบร้อยแล้ว
ไม่ควรทำการพัฒนาโค้ดโดยตรงใน main แต่ให้ทำงานในสาขาอื่นแล้วจึงทำ Pull Request (PR) เพื่อรวมเข้ากับ main
2. development (สาขาพัฒนา)
สาขา development เป็นสาขาที่ใช้สำหรับการรวมโค้ดฟีเจอร์ต่าง ๆ ก่อนจะนำไปทดสอบใน staging
นักพัฒนาควรทำงานในสาขาย่อยของตนเอง (เช่น feature/*) และทำการ PR เพื่อรวมเข้ากับ development
ไม่ควรพัฒนาโค้ดโดยตรงใน development ให้สร้างสาขาย่อยเพื่อพัฒนาแทน
3. stagingTest
สาขานี้สามารถใช้สำหรับการทดสอบการรวมระบบ (integration testing) ก่อนที่จะนำไปยัง main
เป็นสาขาที่รวบรวมการเปลี่ยนแปลงจาก development เมื่อพร้อมที่จะทดสอบและตรวจสอบว่าโค้ดทั้งหมดทำงานได้อย่างถูกต้อง
การทดสอบในสาขา staging จะช่วยให้มั่นใจว่าโค้ดจาก development ไม่มีข้อผิดพลาดก่อนจะดันไป main
4. feature/* (สาขาฟีเจอร์)
สาขา feature/feature และ feature-feature คือสาขาย่อยที่นักพัฒนาสร้างขึ้นเพื่อพัฒนาฟีเจอร์หรือการแก้ไขบางอย่าง

แต่ละคนควรสร้างสาขา feature/* ของตัวเองโดยใช้คำสั่ง

git checkout -b feature/ชื่อฟีเจอร์
ไม่ควรทำงานใน development หรือ main โดยตรง ควรสร้างสาขา feature สำหรับแต่ละฟีเจอร์หรือบั๊กที่จะแก้ไข

แนวทางในการทำ PR และ Merge
การพัฒนาฟีเจอร์ใน feature/*:

แต่ละคนสร้างสาขา feature/* ของตัวเองเพื่อพัฒนาฟีเจอร์ใหม่ ๆ หรือแก้ไขปัญหาต่าง ๆ

เมื่อทำงานเสร็จแล้ว ให้ทำการ commit และ push สาขา feature/* ไปยังรีโมต:

git push origin feature/ชื่อฟีเจอร์
Pull Request (PR) ไปยัง development:

เมื่อพัฒนาเสร็จในสาขา feature/* ให้นักพัฒนาสร้าง PR เพื่อขอรวมโค้ดเข้าสู่ development
การ PR จะช่วยให้สมาชิกทีมคนอื่น ๆ สามารถตรวจสอบโค้ดก่อนที่จะทำการ merge
การรวม (Merge) สาขา development เข้ากับ staging:

หลังจากที่มีการ PR และ merge ฟีเจอร์ทั้งหมดเข้าใน development ให้ทำการทดสอบระบบทั้งหมดใน staging
ถ้าไม่มีปัญหา สามารถ merge เข้าสู่ main ได้
การ PR และ Merge เข้าสู่ main:

เมื่อทดสอบใน staging เสร็จและมั่นใจว่าโค้ดไม่มีปัญหา ให้สร้าง PR เพื่อ merge จาก staging เข้าสู่ main
หลังจาก merge เข้าสู่ main แล้ว การเปลี่ยนแปลงจะพร้อมที่จะปล่อยใช้งานจริง
ตัวอย่างคำสั่งที่ใช้:
สร้างสาขา feature/*:
PR ไปยัง development ผ่าน GitHub เพื่อให้สมาชิกทีมคนอื่นตรวจสอบ

หลังจากตรวจสอบและทดสอบเสร็จ ให้ merge เข้า main ผ่าน PR จาก staging

สรุป:
main: สาขาโค้ดเสถียรสำหรับการใช้งานในโปรดักชัน
development: สาขารวมฟีเจอร์ที่อยู่ระหว่างการพัฒนา
staging: สาขาทดสอบก่อนนำไปใช้ในโปรดักชัน
feature/*: สาขาสำหรับการพัฒนาฟีเจอร์แต่ละอัน