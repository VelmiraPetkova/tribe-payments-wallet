## **Wallet API Documentation**  

## **✅ Summary of Implementation**  
This Go application is an HTTP API for managing user wallets. It uses:  
- **`chi` Router** for handling routes.  
- **`logger.StructuredLogger`** for structured logging.  
- **`render`** for request/response handling.  
- **`wallet.NewWalletRepository()`** as a mock database.  
- **HTTP Methods** for CRUD operations on wallets.  

---

## **🔍 How I Solved the Task**  

### **1️⃣ Structured Routing with `chi`**  
- Registered routes under `/v1/wallet` and grouped related operations.  
- Used dynamic URL parameters (`/wallet/{walletID}`) for wallet-specific operations.  

### **2️⃣ Request Binding and Validation**  
- Created structs (`WalletRequestJson`, `DepositRequestJson`, `WithDrawRequestJson`) to parse JSON requests.  
- Implemented `Bind()` methods for validation.  

### **3️⃣ Wallet Operations**  
- **Create Wallet:** Takes `name` and associates it with a user.  
- **Get Wallet:** Fetches the wallet based on `walletID` and `userID`.  
- **Deposit Money:** Adds a valid amount to the wallet.  
- **Withdraw Money:** Checks if the balance is sufficient before withdrawal.  

---

## **🚀 Suggested Improvements**  

### **1️⃣ Better Error Handling**  
- Instead of `w.WriteHeader(http.StatusBadRequest)`, send **JSON responses** with an error message.  
- Example:  
  ```go
  http.Error(w, `{"error": "Invalid wallet ID"}`, http.StatusBadRequest)
  ```  

### **2️⃣ Use Database Instead of `walletStore`**  
- `walletStore` is like an in-memory store. Consider using **PostgreSQL with GORM**:  
  ```go
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  ```  

  ```  
- Apply it globally:  
  ```go
  mux.Use(AuthMiddleware)
  ```  

---

# **🚀 How to Run the API**  

### **1️⃣ Clone the Repository**  
```sh
git clone https://github.com/your-repo/wallet-api.git
cd wallet-api
```  

### **2️⃣ Install Dependencies**  
```sh
go mod tidy
```  

### **3️⃣ Run the API**  
```sh
go run main.go api
```  

### **4️⃣ API runs on:**  
```
http://localhost:8080
```  

---

## **📌 API Endpoints & Example Requests**  

### **1️⃣ Create a Wallet**  
**Request:**  
```http
POST http://localhost:8080/v1/wallet
```  
**Headers:**  
```
X-Auth-UserId: Mira
Content-Type: application/json
```  
**Body:**  
```json
{
  "name": "Test"
}
```  
**Response:**  
```json
{
  "ID": "c1b5a1ff-4ae0-4064-8f07-53766128f8bf",
  "Balance": 0,
  "Name": "Test",
  "Iban": "43acba62-57d8-4654-aaa2-bd8558a4879e",
  "UserId": "Mira"
}
```  

---

### **2️⃣ Get Wallet Details**  
**Request:**  
```http
GET http://localhost:8080/v1/wallet/{id}
```  
**Headers:**  
```
X-Auth-UserId: Mira
```  
**Response:**  
```json
{
  "ID": "c1b5a1ff-4ae0-4064-8f07-53766128f8bf",
  "Balance": 40,
  "Name": "Test",
  "Iban": "43acba62-57d8-4654-aaa2-bd8558a4879e",
  "UserId": "Mira"
}
```  

---

### **3️⃣ Deposit Money into Wallet**  
**Request:**  
```http
POST http://localhost:8080/v1/wallet/{id}/deposit
```  
**Headers:**  
```
X-Auth-UserId: Mira
Content-Type: application/json
```  
**Body:**  
```json
{
  "depositValue": 50
}
```  
**Response:**  
```json
{
  "ID": "c1b5a1ff-4ae0-4064-8f07-53766128f8bf",
  "Balance": 50,
  "Name": "Test",
  "Iban": "43acba62-57d8-4654-aaa2-bd8558a4879e",
  "UserId": "Mira"
}
```  

---

### **4️⃣ Withdraw Money from Wallet**  
**Request:**  
```http
POST http://localhost:8080/v1/wallet/{id}/withdraw
```  
**Headers:**  
```
X-Auth-UserId: Mira
Content-Type: application/json
```  
**Body:**  
```json
{
  "withDrawValue": 10
}
```  
**Response:**  
```json
{
  "ID": "c1b5a1ff-4ae0-4064-8f07-53766128f8bf",
  "Balance": 40,
  "Name": "Test",
  "Iban": "43acba62-57d8-4654-aaa2-bd8558a4879e",
  "UserId": "Mira"
}
```  

---

## **✅ Notes**  
- `{id}` should be replaced with the actual wallet ID returned from the **Create Wallet** request.  
- The API requires the **X-Auth-UserId** header for authentication.  
- If balance is insufficient during withdrawal, a **400 Bad Request** response is returned.  