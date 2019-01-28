package agouti_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"

//	. "agouti"
)

var _ = Describe("UserLogin", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})
	
	It("should manage user authentication", func() {
        By("redirecting the user to the login form", func() {
            Expect(page.Navigate("http://localhost:3000")).To(Succeed())
            Expect(page).To(HaveURL("http://localhost:3000/login"))
        })

        By("allowing the user to fill out the login form", func() {
//            userField := page.FindByLabel("User")
//            Eventually(userField).Fill("bob").Should(Succeed())
            passwordField := page.FindByLabel("Password")
            Expect(passwordField.Fill("secret")).To(Succeed())
            Expect(page.FindByLabel("Remember Me").Check()).To(Succeed())
            Expect(page.Find("#login_form").Submit()).To(Succeed())
        })

        By("directing the user to the dashboard", func() {
            Eventually(page).Should(HaveTitle("Dashboard"))
        })

        By("allowing the user to view his or her profile", func() {
            Expect(page.FindByLink("Profile Page").Click()).To(Succeed())
            profile := page.Find("section.profile")
            greeting := profile.Find(".greeting")
            Eventually(greeting).Should(HaveText("Hello!"))
            Expect(profile.Find("img#profile_pic")).To(BeVisible())
        })

        By("allowing the user to log out", func() {
            Expect(page.Find("#logout").Click()).To(Succeed())
            Expect(page).To(HavePopupText("Are you sure?"))
            Expect(page.ConfirmPopup()).To(Succeed())
            Eventually(page).Should(HaveTitle("Login"))
        })
    })
})
