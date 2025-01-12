from selenium.webdriver.common.by import By
from selenium.webdriver.firefox.webdriver import WebDriver
from selenium.common.exceptions import NoSuchElementException

def check_element_value_by_name(driver: WebDriver, el_name: str, value: str) -> None:
    el = driver.find_element(By.NAME, el_name)
    assert el.text == value, f"element with name {el_name} text must be equal to {value}, but got {el.text}"

def check_no_element(driver: WebDriver, el_name: str) -> None:
    try:
        driver.find_element(By.NAME, el_name)
    except NoSuchElementException:
        return
    
    raise Exception(f"element with name {el_name} is expected to be None")