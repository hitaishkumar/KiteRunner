import logging
import time
import pyotp
from kiteconnect import KiteConnect, KiteTicker
from datetime import datetime, time as dtime

# --- CONFIGURATION ---
API_KEY = "YOUR_API_KEY"
API_SECRET = "YOUR_API_SECRET"
USER_ID = "YOUR_ZERODHA_USER_ID"
PASSWORD = "YOUR_ZERODHA_PASSWORD"
# TOTP_KEY: Get this from Zerodha > My Profile > Password & Security > Enable TOTP (Copy the text key, don't just scan QR)
TOTP_KEY = "YOUR_TOTP_TEXT_KEY_HERE" 

# Strategy Settings
SYMBOL = "NSE:INFY"         # Instrument to trade
TOKEN = 408065              # Instrument Token (Check Zerodha instrument list)
QUANTITY = 10
ORB_TIME_MINUTES = 15       # 15-minute Opening Range
STOP_LOSS_PCT = 1.0         # 1% SL
TARGET_PCT = 2.0            # 2% Target

# Global Variables
orb_high = 0
orb_low = 0
order_placed = False
kite = None

# --- AUTOMATED LOGIN SYSTEM ---
def autologin():
    global kite
    print("[*] Attempting Auto-Login...")
    kite = KiteConnect(api_key=API_KEY)
    
    # 1. Generate TOTP
    print("[*] Generating TOTP...")
    totp = pyotp.TOTP(TOTP_KEY)
    current_otp = totp.now()
    
    # NOTE: Fully automated login requires using a headless browser (Selenium/Playwright) 
    # to submit USER_ID, PASSWORD, and OTP to get the 'request_token'.
    # Zerodha forbids storing user passwords directly in scripts for full bypass without user interaction 
    # due to regulations. 
    # FOR THIS CODE TO WORK: You must manually get the 'request_token' once a day from the login URL 
    # or implement a Selenium wrapper (advanced).
    
    # Simplified Manual Flow for Compliance:
    print(f"[*] Login URL: {kite.login_url()}")
    request_token = input("[-] Paste the 'request_token' from the redirected URL here: ")
    
    data = kite.generate_session(request_token, api_secret=API_SECRET)
    kite.set_access_token(data["access_token"])
    print(f"[*] Login Successful! Access Token set.")
    return kite

# --- ORDER MANAGEMENT ---
def place_order(transaction_type, price):
    global order_placed
    try:
        # TRIGGER_PRICE used for SL-M or simply MARKET order here for speed
        order_id = kite.place_order(
            tradingsymbol=SYMBOL.split(":")[1],
            exchange=kite.EXCHANGE_NSE,
            transaction_type=transaction_type,
            quantity=QUANTITY,
            variety=kite.VARIETY_REGULAR,
            order_type=kite.ORDER_TYPE_MARKET,
            product=kite.PRODUCT_MIS,
            validity=kite.VALIDITY_DAY
        )
        print(f"✅ ORDER PLACED: {transaction_type} at approx {price}. ID: {order_id}")
        order_placed = True
        # IMMEDIATELY PLACE STOP LOSS & TARGET ORDERS HERE (Omitted for brevity)
    except Exception as e:
        print(f"❌ Order Failed: {e}")

# --- REAL-TIME TICK HANDLER ---
def on_ticks(ws, ticks):
    global orb_high, orb_low, order_placed
    
    # Get current time
    now = datetime.now().time()
    
    for tick in ticks:
        ltp = tick['last_price']
        
        # 1. ORB FORMATION PHASE (e.g., 9:15 - 9:30)
        orb_end_time = dtime(9, 15 + ORB_TIME_MINUTES) # Simplified time addition
        
        if now < orb_end_time:
            if orb_high == 0: orb_high = ltp
            if orb_low == 0: orb_low = ltp
            
            if ltp > orb_high: orb_high = ltp
            if ltp < orb_low: orb_low = ltp
            # print(f"Forming Range... High: {orb_high} Low: {orb_low}")
            
        # 2. TRADING PHASE
        elif not order_placed:
            # print(f"Waiting for Breakout... LTP: {ltp} Range: {orb_low}-{orb_high}")
            
            if ltp > orb_high:
                print(f"🚀 BREAKOUT BUY DETECTED at {ltp}!")
                place_order(kite.TRANSACTION_TYPE_BUY, ltp)
                
            elif ltp < orb_low:
                print(f"🔻 BREAKDOWN SELL DETECTED at {ltp}!")
                place_order(kite.TRANSACTION_TYPE_SELL, ltp)

def on_connect(ws, response):
    print(f"[*] Connected to KiteTicker. Subscribing to {TOKEN}...")
    ws.subscribe([TOKEN])
    ws.set_mode(ws.MODE_FULL, [TOKEN])

def on_close(ws, code, reason):
    print(f"[*] Connection Closed: {reason}")

# --- MAIN EXECUTION ---
if __name__ == "__main__":
    try:
        kite = autologin()
        
        # Initialize WebSocket
        kws = KiteTicker(API_KEY, kite.access_token)
        kws.on_ticks = on_ticks
        kws.on_connect = on_connect
        kws.on_close = on_close
        
        print("[-] Starting Ticker...")
        kws.connect()
        
    except KeyboardInterrupt:
        print("\n[!] Bot stopped by user.")