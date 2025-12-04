# 🦅 KiteRunner — Trade from Your Terminal

KiteRunner is an open-source, terminal-based trading interface built on the **Kite Connect API**, allowing you to trade, monitor markets, and manage your portfolio — all from the CLI.

---

### Login Page

![Login Page](<assets/image%20(2).png>)

### Menu Page

![Menu Page](<assets/image%20(3).png>)

### Home Page

![Home Page](<assets/image%20(4).png>)

### Instruments Page

![Instruments Page](<assets/image%20(5).png>)

### Order History Page

In the Order History menu, you can:

1. View Order Details
2. Order History
3. Order Trades

![Order History Page](<assets/image%20(7).png>)

### Trade Book for an Order

![Trade Book](<assets/image%20(8).png>)

### Order Details Page

![Order Details](<assets/image%20(9).png>)

### Extended Order History

![Extended Order History](<assets/image%20(10).png>)

## Features

### Current Functionality

KiteRunner currently supports the core workflows needed for trading directly from the terminal:

- Interactive CLI interface built using Go and tview
- Secure login and session handling
- Live market data streaming (LTP, OHLC, market depth)
- Ability to place, modify, and cancel orders
- Complete order history and trade book view
- Access to portfolio, positions, holdings, and available funds
- Fast instrument search and filtering
- Concurrent data fetchers to reduce latency

---

## Roadmap / Upcoming Features

These are the planned improvements and additions, based on what the Kite Connect API supports and what I find useful.

### Order Types

- Support for all order varieties:
  - Market, Limit, SL, SL-M
  - Cover Orders (CO)
  - Iceberg Orders
  - After Market Orders (AMO)

### GTT / Conditional Orders

- Create and manage Good-Till-Triggered (GTT) orders
- Trigger-based automated execution

### Historical Data and Backtesting

- Fetch historical candle data (1-minute, 5-minute, 15-minute, end-of-day)
- Basic CLI-based backtesting for simple strategy validation

### Basket Orders

- Create and execute multiple orders as a basket
- Useful for multi-leg strategies and quick portfolio rebalancing

### Postbacks and Webhooks

- Real-time order and trade updates pushed to the terminal or user scripts

### Advanced Portfolio Insights

- Margin breakdown
- Basic P&L summaries
- Optional risk and volatility indicators (depending on API limitations)

### Developer and Power-User Features

- Strategy tags for organizing and tracking orders
- Export of trade history and market data to JSON or CSV
- CLI auto-complete, watchlists, and presets
- More flexible configuration (API keys, timeouts, logging options)

## 🏗️ Built With

- **Golang** (main engine)
- **Cobra** (tview framework)
- **Kite Connect API** (trading + data layer)
- **WebSockets** (market streaming)

---

## 🤝 Contributing

PRs are very welcome!

---
