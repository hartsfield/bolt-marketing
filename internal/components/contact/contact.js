{{ define "contact.js" }}
async function submitContact() {
    var name = document.getElementById("name-input").value;
    var email = document.getElementById("email-input").value;
    var phone = document.getElementById("phone-input").value;
    const response = await fetch("/api/submitContact", {
        method: "POST",
        headers: { 
            "Content-Type": "application/json", 
            'Accept': 'application/json'
        },
        body: JSON.stringify({
            "name": name,
            "email": email,
            "phone": phone,
        }),
    });

    let res = await response.json();
    if (res.success == "true") {
        console.log("success");
    } else {
        console.log("error");
    }
}
{{end}}
