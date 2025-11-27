import random
import time
from collections import deque
from datetime import datetime, time as dtime, timedelta

# ===== STRATEGY SETTINGS =====
SYMBOL = "NSE:INFY"
QUANTITY = 10
ORB_TIME_MINUTES = 15
STOP_LOSS_PCT = 1.0
TARGET_PCT = 2.0
VOLUME_FILTER_FACTOR = 1.5  # breakout valid only if volume > 1.5x avg vol.

# ===== STATE VARIABLES =====
orb_high = 0
orb_low = 0
order_placed = False
side = None
entry_price = None
stop_loss = None
target = None
trade_closed = False

# store last few candle volumes for avg calculation
recent_volumes = deque(maxlen=5)

# ===== PRINT HELPERS =====
def hr(): print("────────────────────────────────────────────")
def log(msg): print(f"[{datetime.now().strftime('%H:%M:%S')}] {msg}")

# ===== ORDER MANAGEMENT =====
def place_order(transaction_type, price):
    global order_placed, side, entry_price, stop_loss, target
    
    side = transaction_type
    entry_price = price
    order_placed = True
    hr()
    log(f"📌 ORDER EXECUTED → {transaction_type} @ {price}")

    # define SL and Target
    if transaction_type == "BUY":
        stop_loss = price * (1 - STOP_LOSS_PCT / 100)
        target = price * (1 + TARGET_PCT / 100)
    else:
        stop_loss = price * (1 + STOP_LOSS_PCT / 100)
        target = price * (1 - TARGET_PCT / 100)

    log(f"🎯 TARGET: {round(target, 2)} | 🛑 STOP LOSS: {round(stop_loss, 2)}")
    hr()

# ===== TRADE EXIT MANAGER =====
def close_trade(exit_price):
    global trade_closed, entry_price, side
    trade_closed = True
    pnl = (exit_price - entry_price) * QUANTITY
    if side == "SELL": pnl = -pnl
    log(f"💵 TRADE CLOSED @ {exit_price} | 📊 P&L: {round(pnl,2)}")
    hr()

# ===== STRATEGY CORE WITH VOLUME FILTER =====
def on_tick(price, volume, now):
    global orb_high, orb_low, order_placed, trade_closed, side, entry_price, stop_loss, target
    
    recent_volumes.append(volume)
    avg_volume = sum(recent_volumes) / len(recent_volumes) if recent_volumes else 0

    orb_end_time = dtime(9, 15 + ORB_TIME_MINUTES)

    # 🏛 ORB FORMATION PHASE
    if now < orb_end_time and not order_placed:
        if orb_high == 0: orb_high = price
        if orb_low == 0: orb_low = price
        if price > orb_high: orb_high = price
        if price < orb_low: orb_low = price
        
        log(f"📌 ORB FORMING → HIGH:{orb_high} | LOW:{orb_low} | LTP:{price} | VOL:{volume}")
        return

    # 🚦 Breakout Checks + Volume Filter
    if not order_placed:
        if price > orb_high:
            if volume > avg_volume * VOLUME_FILTER_FACTOR:
                log(f"🚀 VALID BREAKOUT BUY → LTP:{price} | VOL:{volume} > {round(avg_volume,2)}")
                place_order("BUY", price)
            else:
                log(f"⚠️ LOW VOLUME BREAKOUT ❌ Rejected → LTP:{price} | VOL:{volume} < {round(avg_volume,2)}")

        elif price < orb_low:
            if volume > avg_volume * VOLUME_FILTER_FACTOR:
                log(f"🔻 VALID BREAKDOWN SELL → LTP:{price} | VOL:{volume} > {round(avg_volume,2)}")
                place_order("SELL", price)
            else:
                log(f"⚠️ LOW VOLUME BREAKDOWN ❌ Rejected → LTP:{price} | VOL:{volume} < {round(avg_volume,2)}")
        else:
            log(f"⏳ Waiting | LTP:{price} | Range:{orb_low}-{orb_high} | VOL:{volume}")
        return

    # 📉 Manage SL/Target if trade exists
    if order_placed and not trade_closed:
        if side == "BUY":
            if price <= stop_loss:
                hr(); log(f"💔 STOPLOSS HIT BUY @ {price}")
                close_trade(price)
            elif price >= target:
                hr(); log(f"🎉 TARGET HIT BUY @ {price}")
                close_trade(price)

        elif side == "SELL":
            if price >= stop_loss:
                hr(); log(f"💔 STOPLOSS HIT SELL @ {price}")
                close_trade(price)
            elif price <= target:
                hr(); log(f"🎉 TARGET HIT SELL @ {price}")
                close_trade(price)

# ===== RANDOM TICK SIMULATION =====
def run_simulation(duration_sec=120):
    log(f"🎬 STARTING SIMULATION (Volume + ORB Filter) for {duration_sec} seconds...")
    hr()

    start_price = random.randint(1500, 1600)
    current_price = start_price

    start_time = datetime.now()
    fake_time = datetime.combine(datetime.today(), dtime(9, 16))

    while (datetime.now() - start_time).seconds < duration_sec:
        current_price += random.randint(-6, 6)
        volume = random.randint(50000, 300000)  # realistic NSE volumes

        on_tick(current_price, volume, fake_time.time())

        fake_time += timedelta(seconds=30)
        time.sleep(0.8)

    hr(); log("🏁 SIMULATION ENDED.")

# ======== RUN ==========
if __name__ == "__main__":
    run_simulation(120)
