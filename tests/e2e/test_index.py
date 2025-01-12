from selenium import webdriver
from selenium.webdriver.firefox import service
from selenium.webdriver.common.by import By
import common
import logging

def run() -> None:
    logging.root.setLevel(logging.INFO)
    test_index()

def test_index() -> None:
    try:
        run_index_test()
    except Exception as ex:
        logging.error(f"test_index: error: {ex}")
    else:
        logging.info("test_index: ok")

def run_index_test() -> None:
    gecko_service = service.Service(executable_path="/snap/bin/firefox.geckodriver")
    with webdriver.Firefox(service=gecko_service) as driver:
        driver.get("http://localhost:8000/")

        common.check_element_value_by_name(driver, "a", "")
        common.check_element_value_by_name(driver, "b", "")
        common.check_no_element(driver, "error")
        common.check_no_element(driver, "prev")
        common.check_no_element(driver, "result")
