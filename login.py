from playwright.sync_api import sync_playwright
import json

def save_cookies():
    with sync_playwright() as p:
        browser = p.chromium.launch(headless=False)
        context = browser.new_context()
        page = context.new_page()

        print("Opening Codeforces login page...")
        page.goto("https://codeforces.com/enter")

        print("Please log in and solve CAPTCHA if needed...")
        page.wait_for_timeout(30000)  # Wait for user to login + CAPTCHA

        cookies = context.cookies()
        with open("cookies.json", "w") as f:
            json.dump(cookies, f, indent=2)

        print("✅ Cookies saved to cookies.json")
        browser.close()

if __name__ == "__main__":
    save_cookies()
