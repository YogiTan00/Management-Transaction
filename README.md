# 📊 Management Transaction

A comprehensive transaction management system for tracking daily, weekly, monthly, and yearly transactions. Built with **Go** (Golang) for high-performance backend processing.

## 📋 Overview

This application provides robust transaction management capabilities with features including:
- Daily, weekly, monthly, and yearly transaction tracking
- Product management with flexible pricing and discount options
- Transaction history with return handling
- Soft delete functionality for data integrity

---

## 🗂️ Data Entities

### Product Entity

| Field | Type | Description |
|-------|------|-------------|
| `ID` | UUID/Int | Unique product identifier |
| `Name` | String | Product name |
| `Type` | String | Book type (A5 or B5) |
| `Price` | Decimal | Normal product price |
| `Discount` | Int | Discount percentage (0-100%) |
| `DiscountPrice` | Decimal | Auto or manual discount price |
| `CreatedAt` | Timestamp | Record creation time (auto-generated) |
| `UpdatedAt` | Timestamp | Last update time (auto-generated) |
| `DeletedAt` | Timestamp | Soft delete timestamp |

### Transaction Entity

| Field        | Type | Description |
|--------------|------|-------------|
| `ID`         | UUID/Int | Unique transaction identifier |
| `IDProduct`  | UUID/Int | Reference to product ID |
| `UnitReturn` | Int | Number of returned units (default: 0) |
| `UnitTotal`  | Int | Number of units sold |
| `Date`       | Date | Transaction date |
| `Notes`      | String | Additional transaction notes |
| `CreatedAt`  | Timestamp | Record creation time (auto-generated) |
| `UpdatedAt`  | Timestamp | Last update time (auto-generated) |
| `DeletedAt`  | Timestamp | Soft delete timestamp |

---

## 🛠️ Tech Stack & Infrastructure

This project is built with modern Go ecosystem tools:

- **[Gorm](https://gorm.io/)** - The fantastic ORM library for Golang
- **[Gin](https://gin-gonic.com/)** - High-performance HTTP web framework

---

## 🚀 Getting Started

### Prerequisites
- Go 1.25 or higher
- Database (PostgreSQL/MySQL/SQLite)

### Installation
```bash
# Clone the repository
git clone https://github.com/yourusername/Management-Transaction.git

# Navigate to project directory
cd Management-Transaction

# Install dependencies
go mod download

# Run the application
go run main.go
```

---

## 📝 License

This project is licensed under the GNU Affero General Public License v3.0 (AGPL-3.0) License.
