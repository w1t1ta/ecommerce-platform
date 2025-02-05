FROM postgres:14-alpine

# คัดลอกไฟล์ init.sql ไปยังตำแหน่งที่ PostgreSQL จะทำการรัน script อัตโนมัติ
COPY init.sql /docker-entrypoint-initdb.d/

# เปิดพอร์ต 5432 สำหรับ PostgreSQL
EXPOSE 5432