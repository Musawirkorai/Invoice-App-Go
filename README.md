# 📄 Go Invoice System

A **web-based invoice system** built with **Go (Golang)** for the backend and **HTML/CSS + JavaScript** for the frontend.
Features:

* Add multiple items dynamically
* Automatic calculation of item totals and grand total
* Customer name and invoice number input
* Generate invoice in a professional format in a new tab
* Inline CSS with a clean, modern UI

---

## ⚙️ Features

* Add multiple items dynamically
* Auto-calculated totals (per item and total bill)
* Customer name and optional invoice number
* Professional invoice table with date and store info
* Beautiful UI with alternating row colors, bold totals, rounded buttons
* No database required (all in-memory)

---

## 🖥️ Getting Started (Local)

### **Prerequisites**

* Go installed (version 1.20+ recommended)

### **Steps**

1. Clone the repository:

```bash
git clone https://github.com/yourusername/invoice-app.git
cd invoice-app
```

2. Run the app:

```bash
go run main.go
```

3. Open the app in your browser:

```
http://localhost:8080
```

4. Add customer name, items, and click **Generate Invoice**. The invoice will open in a new tab.

---

## 🚀 Deploying on AWS EC2

Follow these steps to deploy your invoice system online:

### **1. Create an EC2 Instance**

* AWS Console → EC2 → Launch Instance
* Choose **Amazon Linux 2** or **Ubuntu 22.04**
* Free tier: **t2.micro**
* Configure SSH key pair

---

### **2. Connect to EC2**

```bash
ssh -i "your-key.pem" ec2-user@<EC2-Public-IP>
```

*(Ubuntu AMI: `ssh -i "your-key.pem" ubuntu@<EC2-Public-IP>`)*

---

### **3. Install Go**

**Amazon Linux:**

```bash
sudo yum update -y
sudo yum install golang -y
```

**Ubuntu:**

```bash
sudo apt update
sudo apt install golang-go -y
```

Check Go version:

```bash
go version
```

---

### **4. Upload Your Project**

**Option A:** Clone from GitHub

```bash
git clone https://github.com/yourusername/invoice-app.git
cd invoice-app
```

**Option B:** Upload manually using WinSCP, FileZilla, or S3

---

### **5. Run Your Go App**

```bash
go run main.go
```

* Opens on **[http://localhost:8080](http://localhost:8080)**
* Accessible globally once port 8080 is opened

---

### **6. Allow Port 8080 in Security Group**

* EC2 → Security Group → Edit inbound rules → Add:

  * Type: Custom TCP
  * Port: 8080
  * Source: 0.0.0.0/0

Now open:

```
http://<EC2-Public-IP>:8080
```

---

### **7. Optional: Run in Background**

```bash
nohup go run main.go > output.log 2>&1 &
```

---

### **8. Optional: Use Nginx Reverse Proxy**

* Install Nginx:

```bash
sudo yum install nginx -y   # Amazon Linux
sudo apt install nginx -y   # Ubuntu
sudo systemctl start nginx
```

* Proxy requests to Go app (`/etc/nginx/nginx.conf`):

```nginx
location / {
    proxy_pass http://127.0.0.1:8080;
}
```

* Restart Nginx:

```bash
sudo systemctl restart nginx
```

* Access app via **http://<EC2-Public-IP>** (port 80)

---

## 📦 Project Structure

```
invoice-app/
│
├── main.go          # Go backend
├── README.md        # Project documentation
└── go.mod           # Go modules
```

---

## 🌟 Future Enhancements

* Printable PDF invoice download
* Dark mode UI
* Auto-save invoices in memory or database
* Email invoices to customers

---

## 📝 Author

Musawir Ali

* GitHub: Musawirkorai (https://github.com/Musawirkorai)

