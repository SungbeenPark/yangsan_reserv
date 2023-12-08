package main

import (
	"github.com/fedesog/webdriver"
	"github.com/tebeka/selenium"
	"log"
	"time"
)

func main() {
	chromeDriver := webdriver.NewChromeDriver("./chromedriver")
	err := chromeDriver.Start()
	if err != nil {
		log.Println("chromeDriver Start Err ", err)
	} else {
		desired := webdriver.Capabilities{"Platform": "Windows"}
		required := webdriver.Capabilities{}

		session, err := chromeDriver.NewSession(desired, required)
		if err != nil {
			log.Println("Session Err : ", err)
		}
		err = session.Url("https://hscamping.yssisul.or.kr:453/rsvc/rsv_srm.html?b_id=hscamping")
		if err != nil {
			log.Println(err)
		}

		log.Println("session start success")
		// 로그인 클릭
		btn, _ := session.FindElement(selenium.ByXPATH, "/html/body/div[2]/div[1]/div/span[4]/a")
		err = btn.Click()
		if err != nil {
			log.Println("click err ", err)
		}

		//*[@id="m_email"]
		//로그인 계정정보 입력
		idInput, err := session.FindElement(selenium.ByID, "m_email")
		if err != nil {
			log.Println("idInput err ", err)
		}
		idInput.SendKeys("been778")

		pwInput, err := session.FindElement(selenium.ByID, "m_pwdTmp")
		if err != nil {
			log.Println("pwInput err ", err)
		}
		pwInput.SendKeys("dlfdltka1!")

		loginBtn, err := session.FindElement(selenium.ByXPATH, "//*[@id=\"main_layout\"]/div/div/div[2]/div/div[1]/form/button")
		if err != nil {
			log.Println("loginBtn err ", err)
		}
		loginBtn.Click()
		yangsanBtn, err := session.FindElement(selenium.ByXPATH, "//*[@id=\"form1\"]/div[1]/div[2]/div/div/div/ul/li[2]/span/b")
		if err != nil {
			log.Println("yangsanBtn err ", err)
		}
		yangsanBtn.Click()

		//agreeUseAuth
		agreeUseAuth, err := session.FindElement(selenium.ByID, "agreeUseAuth")
		if err != nil {
			log.Println("agreeUseAuth err ", err)
		}
		agreeUseAuth.Click()
		//m2_ssn1
		m2_ssn1, err := session.FindElement(selenium.ByID, "m2_ssn1")
		if err != nil {
			log.Println("m2_ssn1 err ", err)
		}
		m2_ssn1.SendKeys("860301")
		//m2_ssn2
		m2_ssn2, err := session.FindElement(selenium.ByID, "m2_ssn2")
		if err != nil {
			log.Println("m2_ssn2 err ", err)
		}
		m2_ssn2.SendKeys("1120620")
		//useAuthBtn
		useAuthBtn, err := session.FindElement(selenium.ByID, "useAuthBtn")
		if err != nil {
			log.Println("useAuthBtn err ", err)
		}
		useAuthBtn.Click()

		//click Date
		selectDate, err := session.FindElement(selenium.ByID, "20230714")
		if err != nil {
			log.Println("selectDate err ", err)
		}
		selectDate.Click()

		var findResult = true
		for findResult {
			//iframe 찾기
			iframe, err := session.FindElement(selenium.ByXPATH, "/html/body/div[2]/div[2]/div[2]/form/div[3]/div[2]/div[1]/div/div/span/div/div/iframe")
			if err != nil {
				log.Println("iframe err ", err)
				continue
			}

			session.FocusOnFrame(iframe)

			//captcha 내용 확인
			captcha, err := session.FindElement(selenium.ByID, "recaptcha-anchor")
			if err != nil {
				log.Println("captcha err ", err)
				continue
			}
			captcha.Click()
			findResult = false
		}
		time.Sleep(60 * time.Second * 2)
		session.Delete()
	}

	chromeDriver.Stop()
}
