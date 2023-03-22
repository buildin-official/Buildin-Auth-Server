package utils

import (
	"crypto/tls"
	"fmt"
	"strings"

	"gopkg.in/gomail.v2"
	"pentag.kr/BuildinAuth/configs"
)

const (
	MIME             = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	VerificatonEmail = `<!DOCTYPE html><html lang="ko"> <head> <meta charset="UTF-8"/> <meta name="viewport" content="width=device-width, initial-scale=1.0"/> <title>Clast Email Verify</title> <link href="https://cdn.jsdelivr.net/gh/sunn-us/SUIT/fonts/static/woff2/SUIT.css" rel="stylesheet"/> </head> <body style=" margin: 0; padding: 0; width: 100%; height: 100%; font-size: 16px; font-family: 'SUIT', sans-serif; cursor: default; font-weight: 600; box-sizing: border-box; word-break: keep-all;" > <div style=" padding: 2rem; display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start; " class="wrapper" > <svg xmlns="http://www.w3.org/2000/svg" width="141" height="25" viewBox="0 0 141 25" fill="none" > <path d="M13.5903 12.2375L10.4137 12.1287C10.822 8.65256 13.5378 6.8292 17.0624 6.85052C21.0598 6.8292 24 10.0877 24 14.4541C24 19.3804 21.0169 22.4 16.9766 22.4C13.0866 22.4 10.7535 20.1957 10.4096 17.4021C11.9691 16.7607 12.4792 16.0676 13.4319 14.4033C13.337 15.5572 13.2272 15.8826 13.4319 17.0689C13.7327 18.6683 15.1713 19.6065 16.8906 19.6278C19.2332 19.6065 20.7978 17.717 20.7763 14.4541C20.7978 11.8365 19.3405 9.6014 16.9764 9.62272C15.2142 9.6014 13.8268 10.5528 13.5903 12.2375Z" fill="url(#paint0_linear_807_30096)"/> <path d="M10.4137 12.1287L13.5903 12.2375C13.182 15.7136 10.4622 17.5708 6.93759 17.5495C2.94021 17.5708 0 14.3123 0 9.94591C0 5.01962 2.98306 2 7.02343 2C10.9134 2 13.2465 4.20427 13.5904 6.99794C12.0309 7.63929 11.5208 8.33244 10.5681 9.9967C10.663 8.84279 10.7728 8.51739 10.5681 7.33114C10.2673 5.73169 8.8287 4.79352 7.10939 4.7722C4.76684 4.79352 3.2022 6.68304 3.22369 9.94591C3.2022 12.5635 4.66357 14.7648 7.02761 14.7434C8.7899 14.7648 10.1773 13.8134 10.4137 12.1287Z" fill="url(#paint1_linear_807_30096)"/> <path d="M36.48 20.52C38.8 20.56 40.42 19.64 41.76 18.26L39.68 16.24C39.1 17.04 38.06 17.88 36.46 17.88C33.74 17.88 32.04 15.92 32.04 12.78C32.04 9.62 33.74 7.66 36.46 7.68C37.82 7.68 39.12 8.4 39.82 9.44L41.76 7.3C40.42 5.92 38.56 5.04 36.48 5.04C32.08 5.02 29 8.22 29 12.78C29 17.32 32.08 20.52 36.48 20.52ZM43.355 5.2V20.36H46.395V5.2H43.355ZM53.2244 18C51.7444 18 50.8044 16.82 50.8044 14.94C50.8044 13.06 51.7444 11.88 53.2244 11.88C54.6644 11.88 55.8244 13.22 55.8444 14.94C55.8244 16.66 54.6644 18 53.2244 18ZM47.7644 14.94C47.7644 18.14 49.8844 20.4 52.9044 20.4C54.0044 20.4 55.0844 19.92 55.8444 19.06V20.38L58.8844 20.36V9.46H55.8444V10.84C55.0844 9.96 54.0044 9.48 52.9044 9.48C49.8844 9.48 47.7644 11.74 47.7644 14.94ZM62.735 16.88L60.755 18.52C61.715 19.64 63.415 20.46 64.855 20.52C67.515 20.6 69.415 19.22 69.495 17.12C69.615 13.7 64.075 13.8 64.115 12.68C64.135 12.18 64.615 11.86 65.315 11.88C65.895 11.9 66.635 12.32 67.115 12.88L69.055 11.22C68.215 10.22 66.695 9.5 65.395 9.44C62.915 9.36 61.135 10.68 61.075 12.66C60.955 16.06 66.495 16.1 66.455 17.08C66.435 17.7 65.795 18.1 64.935 18.08C64.235 18.06 63.315 17.56 62.735 16.88ZM78.3373 19.76L77.2773 17.44C77.2773 17.44 76.6773 17.86 76.0773 17.8C75.6173 17.74 75.2973 17.24 75.2973 16.6V12.26H78.0373V9.66H75.2973V6.46H72.2573V9.66H70.2773V12.26H72.2573V16.6C72.2573 18.84 73.6773 20.4 75.7173 20.4C77.2573 20.4 77.7373 20.04 78.3373 19.76ZM92.1245 20.52C94.4445 20.56 96.0645 19.64 97.4045 18.26L95.3245 16.24C94.7445 17.04 93.7045 17.88 92.1045 17.88C89.3845 17.88 87.6845 15.92 87.6845 12.78C87.6845 9.62 89.3845 7.66 92.1045 7.68C93.4645 7.68 94.7645 8.4 95.4645 9.44L97.4045 7.3C96.0645 5.92 94.2045 5.04 92.1245 5.04C87.7245 5.02 84.6445 8.22 84.6445 12.78C84.6445 17.32 87.7245 20.52 92.1245 20.52ZM98.9995 5.2V20.36H102.04V5.2H98.9995ZM108.749 17.98C107.329 17.98 106.449 16.82 106.449 14.96C106.449 13.1 107.329 11.92 108.749 11.92C110.169 11.92 111.049 13.1 111.049 14.96C111.049 16.82 110.169 17.98 108.749 17.98ZM103.409 14.96C103.409 18.16 105.609 20.42 108.749 20.42C111.889 20.42 114.089 18.16 114.089 14.96C114.089 11.74 111.889 9.48 108.749 9.48C105.609 9.48 103.409 11.74 103.409 14.96ZM118.9 15.24V9.46H115.86V15.24C115.86 18.36 117.88 20.34 120.76 20.34C121.84 20.34 122.98 19.62 123.5 18.62V20.36H126.54V9.46H123.5V15.04C123.56 16.7 122.64 17.9 121.28 17.92C119.88 17.92 118.9 16.94 118.9 15.24ZM133.432 20.38C133.772 20.38 135.392 20.2 136.392 19.06V20.38H139.432V5.2H136.392V10.8C135.392 9.48 133.772 9.44 133.432 9.44C130.412 9.44 128.272 11.7 128.272 14.92C128.272 18.12 130.412 20.38 133.432 20.38ZM131.312 15C131.312 13.1 132.232 11.92 133.712 11.92C135.212 11.92 136.392 13.28 136.392 15C136.392 16.72 135.212 18.06 133.712 18.06C132.232 18.06 131.312 16.88 131.312 15Z" fill="black"/> <defs> <linearGradient id="paint0_linear_807_30096" x1="12" y1="2" x2="12" y2="22.4" gradientUnits="userSpaceOnUse" > <stop stop-color="#8020FA"/> <stop offset="1" stop-color="#562CFA"/> </linearGradient> <linearGradient id="paint1_linear_807_30096" x1="12" y1="2" x2="12" y2="22.4" gradientUnits="userSpaceOnUse" > <stop stop-color="#8020FA"/> <stop offset="1" stop-color="#562CFA"/> </linearGradient> </defs> </svg> <h1 style="color: #262626; font-weight: 700">Clast Cloud 이메일 인증</h1> <p style="margin: 0; color: #4c4c4c;">아래의 '이메일 인증' 버튼을 눌러 회원가입을 완료해주세요.</p><a style=" margin: 1.4rem 0; border: none; outline: none; background: #542bfa; color: #fff; padding: 0.8rem 1.2rem; border-radius: 12px; font-size: 1rem; font-weight: 700; cursor: pointer; text-decoration: none; " href="https://clast.kr/console/user/signup/verify/{code}" >이메일 인증</a > <p style="margin: 0; color: #4c4c4c;">이 메일을 <strong style="font-weight: 700; color: #262626;">중요메일</strong>로 설정해주세요.</p><p style="margin: 0; color: #4c4c4c;">사용량 경고, 중요 공지 등이 해당 메일을 통해 발신됩니다.</p><p style="margin: 0; color: #4c4c4c;"> 메일이 스팸으로 처리되지 않도록 <strong style="font-weight: 700; color: #262626;">꼭</strong> 중요 메일로 설정 부탁드려요. </p><hr style="margin: 2rem 0; width: 100%; border: 1px solid #d9d9d9"/> <div style="display: flex; flex-direction: column" class="footer"> <span style="margin-bottom: 0.3rem; font-size: 1.2rem; font-weight: 700; color: #737373;">클라스트</span> <p style="margin: 0.2rem 0; font-size: 0.9rem; color: #999;">사업자 등록번호: 205-21-72138 | 대표: 함준형</p><p style="margin: 0.2rem 0; font-size: 0.9rem; color: #999;">통신판매업 신고번호: 2023-대전서구-0206</p><p style="margin: 0.2rem 0; font-size: 0.9rem; color: #999;">35344 대전광역시 서구 배재로197번길 27-46</p><span style="font-size: 1rem; margin-top: 0.6rem; color: #999;">※ 이 메일은 발신 전용입니다.</span> </div></div></body></html>`
)

var emailConf = configs.Config.SMTP

// var addr = emailConf.Host + ":" + strconv.Itoa(emailConf.Port)
var tlsconfig = &tls.Config{
	InsecureSkipVerify: true,
	ServerName:         emailConf.Host,
}

type Mail struct {
	To      string
	Subject string
	Body    string
}

func (mail *Mail) SendEmail() (err error) {
	m := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
	m.SetHeader("From", m.FormatAddress(emailConf.Sender, emailConf.SenderName))
	m.SetHeader("To", mail.To)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", mail.Body)

	d := gomail.NewDialer(emailConf.Host, emailConf.Port, emailConf.User, emailConf.Pass)
	d.TLSConfig = tlsconfig
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println(err)
	return err

}

func (mail *Mail) SendVerificationEmail(to string, code string) (err error) {
	mail.Subject = "BuildIn Email Verification - 빌드인 인증메일"
	mail.Body = strings.Replace(VerificatonEmail, "{code}", code, 1)
	mail.To = to
	return mail.SendEmail()
}
