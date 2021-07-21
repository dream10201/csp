const xhr = new XMLHttpRequest();
document.getElementById("submit").onclick = function () {
    document.getElementById("submit").disabled = "disabled";
    const user = document.getElementById("user").value;
    const old = document.getElementById("old").value;
    const newPwd = document.getElementById("new").value;
    const newPwd_check = document.getElementById("new_check").value;
    if (user.length <= 0 || old.length <= 0 || newPwd.length <= 0 || newPwd_check.length <= 0 || newPwd !== newPwd_check) {
        Swal.fire({
            icon: 'warning',
            text: 'Please check your input!'
        });
        document.getElementById("submit").disabled = "";
        return false;
    }
    xhr.open("POST", "/apply", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send("user=" + user + "&pwd=" + old + "&newPwd=" + newPwd);
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                const result = JSON.parse(xhr.responseText);
                if (!result["status"]) {
                    Swal.fire({
                        icon: 'error',
                        title: 'Oops...',
                        text: 'User does not exist or password is wrong!'
                    });
                } else {
                    Swal.fire({
                        icon: 'success',
                        text: 'Successfully.'
                    });
                }
            } else {
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: "Network Error(" + xhr.status + ")."
                });
            }
            document.getElementById("submit").disabled = "";
        }
    };
}