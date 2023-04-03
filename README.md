# go-calculate-api
go calculate api with higher order function

# lib
- github.com/go-redis/redis/v9
- github.com/gofiber/fiber/v2

# docker
- docker-compose up
- docker-compose down

# โครงสร้างของโปรเจคนี้จะแบ่งเป็นสองส่วนคือ 
1. ส่วนของ Service Calculator
2. ส่วของ Redis
3. แตกต่างจาก repo calculator_api ตรงที่เขียนเป็น closure function
Service Calulator ของเรานั้นจะทำการเชื่อต่อกับ Redis ซึ่งจะอธิบายที่ละ Step ดังนี้:
- สร้าง ConfigMap ของ Redis ขึ้นมาซึ่งเราจะเอา config ตัวนี้ไปใช้กลับ Redis
- ทำการ Deploy Redis ขึ้นไปบน Kubernetes ก่อนพร้อมกับทำ Service Type=ClusterIP
- source code ของเราจะมีส่วนที่ทำการติดต่อกับ Redis ซึ่งต้องเป็น Redis ที่ Deploy อยู่บน Kubernetes
- ในที่นี้เราจะทำแบบไม่ Best Practise เท่าไหร่เราจะไปเอา Cluster IP ของ Service Redis มาโดยตรง
- พิมพ์ Kubectl get service -n test-dev ซึ่ง test-dev เป็น namespace ที่เราทำการ Deploy ลงไป
- เอา ClusterIP และ Port ที่ได้ไปใส่ใน Addr ที่จะติดต่อกับ Redis ใน code
- ทำการ Build Docker image ขึ้นมาซึ่งในที่นี้สร้างเป็น calculator:go-dev
- พิมพ์ docker build -t calculator:go-dev .
- เมื่อได้ images ของ calculator แล้วจะทำการ Deploy ลง Kubernetes
- ทำการ Deploy Calculator ลงไปใน Kubenetes โดยมี Service Type=NodePort กำหนด nodePort เป็น 30144

<p align="center">
  <img src="./images/git_k8s.jpg">
</p>

# curl 
- curl localhost:30144/calculate -X GET -H 'content-type:application/json' -d '{"operator": "+","num1": 5,"num2": 5}'
