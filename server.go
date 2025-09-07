package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// หน้าแรก: ฝัง PDF ด้วย <object>
	app.Get("/", func(c *fiber.Ctx) error {
		html := `<!doctype html>
<html lang="th">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width,initial-scale=1" />
  <title>Thesis ADDMS</title>
  <style>
    html, body {
      height:100%;
      margin:0;
      background:#0b1220;
      font-family: ui-sans-serif,system-ui;
    }
    .viewer {
      width: 100%;
      height: 100%;
      border: none;
      background: #101826;
    }
  </style>
</head>
<body>
  <!-- ให้ object ขยายเต็มทั้ง viewport -->
  <object class="viewer" data="/pdf" type="application/pdf">
    <p>เบราว์เซอร์ของคุณไม่รองรับการแสดง PDF แบบฝัง
       <a href="/raw-pdf" target="_blank" rel="noopener">คลิกเพื่อดาวน์โหลด/เปิดไฟล์</a>
    </p>
  </object>
</body>
</html>`
		return c.Type("html").SendString(html)
	})

	// เสิร์ฟ PDF แบบ inline สำหรับการฝังในหน้าเว็บ (Content-Disposition: inline)
	app.Get("/pdf", func(c *fiber.Ctx) error {
		// Fiber จะตั้ง Content-Type ให้ตามนามสกุลไฟล์โดยอัตโนมัติเมื่อใช้ SendFile
		// แต่เราบังคับเป็น inline เพื่อให้แสดงในหน้า (ไม่บังคับดาวน์โหลด)
		c.Set("Content-Disposition", `inline; filename="ADDMS.pdf"`)
		return c.SendFile("./assets/ADDMS.pdf", false)
	})

	// เสิร์ฟไฟล์ PDF ตรงๆ (เหมาะให้ผู้ใช้เปิดแท็บใหม่/ดาวน์โหลด)
	app.Get("/raw-pdf", func(c *fiber.Ctx) error {
		c.Set("Content-Disposition", `inline; filename="ADDMS.pdf"`)
		return c.SendFile("./assets/ADDMS.pdf", false)
	})

	log.Println("server started: http://localhost:5666")
	log.Fatal(app.Listen(":5666"))
}
