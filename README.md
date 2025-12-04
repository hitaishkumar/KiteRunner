# 🦅 KiteRunner — Trade from Your Terminal

KiteRunner is an open-source, terminal-based trading interface built on the **Kite Connect API**, allowing you to trade, monitor markets, and manage your portfolio — all from the CLI.

---

## 📸 Screenshots

### 🔐 Login Page

![Login Page](<assets/image%20(2).png>)

### 🧭 Menu Page

![Menu Page](<assets/image%20(3).png>)

### 🏠 Home Page

![Home Page](<assets/image%20(4).png>)

### 📊 Instruments Page

![Instruments Page](<assets/image%20(5).png>)

### 📜 Order History Page

In the Order History menu, you can:

1. View Order Details
2. Order History
3. Order Trades

![Order History Page](<assets/image%20(7).png>)

### 💼 Trade Book for an Order

![Trade Book](<assets/image%20(8).png>)

### 📄 Order Details Page

![Order Details](<assets/image%20(9).png>)

### 📜 Extended Order History

![Extended Order History](<assets/image%20(10).png>)

---

## 🚀 Features

### ✅ **Current Functionality**

- 🧰 Fully interactive **CLI UI** using Go + Cobra
- 🔑 Secure login/session handling
- 📈 Live market data streaming (LTP, OHLC, depth)
- 🛒 Order placement, modification, cancellation
- 📜 Full order history + trade book view
- 💼 Portfolio, positions, holdings, funds info
- 🔍 Instrument search + filtering
- 🧵 Concurrent data fetchers for low-latency performance
- 🧱 Clean modular code structure (`auth/`, `orders/`, `quotes/`, `ws/`, `cmd/`, etc.)

---

## 🔮 Roadmap / Upcoming Features

Based on capabilities supported by the **Kite Connect API**, these features are planned:

### 📦 **Order Types**

- Support for _all_ order varieties:
  - Market, Limit, SL, SL-M
  - Cover Orders (CO)
  - Iceberg Orders
  - After Market Orders (AMO)

### 🎯 **GTT / Conditional Orders**

- Create, view, and manage Good-Till-Triggered (GTT) orders
- Automate trade execution based on price triggers

### 📉 **Historical Data + Backtesting**

- Fetch historical candle data (1m, 5m, 15m, EOD)
- Run CLI-based backtesting for quick strategy checks

### 🧺 **Basket Orders**

- Create and execute multiple orders as a single basket
- Perfect for multi-leg strategies and rebalancing

### 🔔 **Postbacks / Webhooks**

- Real-time order/trade notifications pushed to your CLI or custom scripts

### 💸 **Advanced Portfolio Insights**

- Margin breakdown
- P&L summaries
- Risk/volatility estimates (if supported by API limits)

### ⚙️ **Developer/Power User Features**

- Strategy tags for orders
- JSON/CSV export of history, trades, and live ticks
- CLI auto-complete, watchlists, presets
- Functional options for clean configuration (API key, timeouts, logging)

---

## 🏗️ Built With

- **Golang** (main engine)
- **Cobra** (CLI framework)
- **Kite Connect API** (trading + data layer)
- **WebSockets** (market streaming)

---

## 📂 Project Link

👉 https://github.com/hitaishkumar/KiteRunner

---

## 🤝 Contributing

PRs are welcome! Feel free to open issues, suggest new features, or help improve documentation.

---

## 📜 License

Apache-2.0 License

---
