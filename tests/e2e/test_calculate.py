from selenium import webdriver
from selenium.webdriver.firefox import service
from selenium.webdriver.common.by import By
import common
import logging

def run() -> None:
    logging.root.setLevel(logging.INFO)
    test_calculate()

def test_calculate() -> None:
    try:
        run_calculate_test()
    except Exception as ex:
        logging.error(f"test_calculate: error: {ex}")
    else:
        logging.info("test_calculate: ok")

def run_calculate_test() -> None:
    gecko_service = service.Service(executable_path="/snap/bin/firefox.geckodriver")
    with webdriver.Firefox(service=gecko_service) as driver:
        driver.get("http://localhost:8000/")

        tests = [
            {"prev": None,"error": None,"result": "75","a": "25","b": "50","op": "+"},
            {"prev": "Предыдущее значение: 75","error": None,"result": "100","a": "20","b": "5","op": "*"},
            {"prev": "Предыдущее значение: 100","error": None,"result": "100","a": "20","b": "5","op": "*"},
            {"prev": "Предыдущее значение: 100","error": "Ошибка: overflow","result": None,"a": "99999999999","b": "9999999999","op": "*"},
            {"prev": "Предыдущее значение: 100","error": None,"result": "20","a": "100","b": "5","op": "/"},
            {"prev": "Предыдущее значение: 20","error": None,"result": "5","a": "25","b": "20","op": "%"},
            {"prev": "Предыдущее значение: 5","error": "Ошибка: cannot divide by 0","result": None,"a": "25","b": "0","op": "/"},
        ]
        eqs = ["prev", "error", "result"]


        for test in tests:
            input_a = driver.find_element(By.NAME, "a")
            input_b = driver.find_element(By.NAME, "b")
            operation_selector = driver.find_element(By.NAME, "op")
            submit = driver.find_element(By.XPATH, "//input[@type='submit']")

            input_a.clear()
            input_a.send_keys(test["a"])

            input_b.clear()
            input_b.send_keys(test["b"])
            operation_selector.send_keys(test["op"])
            submit.click()

            target_url = "http://localhost:8000/calculate"
            assert driver.current_url == target_url, f"url expected to be {target_url}, but got {driver.current_url}"

            for el_name in ["a", "b", "op"]:
                el = driver.find_element(By.NAME, el_name)
                got = el.get_attribute('value')
                want = test[el_name]
                assert got == want, f'expected {el_name} after submit to be {want}, but got {got}'

            for eq in eqs:
                want = test[eq]
                if want is not None:
                    got = driver.find_element(By.NAME, eq)
                    assert got.text == want, f"wanted {want}, got {got.text}"
                else:
                    common.check_no_element(driver, eq)






