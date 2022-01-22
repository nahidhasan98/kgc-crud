console.log("script linked successfully");

$(document).ready(function () {
    // student list display function
    getStudentList();

    // login form submit function
    $('#loginModal form').on('submit', function () {
        loginRequest();
        return false;
    });

    // add form submit function
    $('#addModal form').on('submit', function () {
        addRequest();
        return false;
    });

    // delete form submit function
    $('#deleteModal form').on('submit', function () {
        deleteRequest();
        return false;
    });

    // error field clearing
    errFieldClearing();

    // add button click function
    $('#add, #addPlus').click(function () {
        $('#addModalLabel').text("Add a Student");
        $('#btnNo').text("Close");
        $('#addSubmit').val("Add");
        $('#addMessage').text("");
        clearAllErrorText();

        $('#sid').text("");
        $('#name').val("");
        $('#session').val("");
        $('#reg').val("");
        $('#roll').val("");
        $('#passingYear').val("");
        $('#phone').val("");
        $('#email').val("");
    });

    // search button click function
    $('#searchBtn').click(function () {
        let imageSrc = $(this.children[0]).attr('src');
        let strArr = imageSrc.split("/");
        let image = strArr[strArr.length - 1];

        if (image == "search.png") {
            $(this.children[0]).attr('src', "../assets/images/hide.png")
            $('#searchRow').css('display', 'table-row');
        } else {
            $(this.children[0]).attr('src', "../assets/images/search.png")
            $('#searchRow').css('display', 'none');
        }
    });

    // search/filter options
    filter();
});

// keyup filtering
function filter() {
    $('#nameSearch').keyup(function () {
        let filterKey = $('#nameSearch').val().toUpperCase();
        let colNum = 3;
        doFilter(filterKey, colNum);
    });

    $('#sessionSearch').keyup(function () {
        let filterKey = $('#sessionSearch').val().toUpperCase();
        let colNum = 4;
        doFilter(filterKey, colNum);
    });

    $('#regSearch').keyup(function () {
        let filterKey = $('#regSearch').val().toUpperCase();
        let colNum = 5;
        doFilter(filterKey, colNum);
    });

    $('#rollSearch').keyup(function () {
        let filterKey = $('#rollSearch').val().toUpperCase();
        let colNum = 6;
        doFilter(filterKey, colNum);
    });

    $('#yearSearch').keyup(function () {
        let filterKey = $('#yearSearch').val().toUpperCase();
        let colNum = 7;
        doFilter(filterKey, colNum);
    });

    $('#phoneSearch').keyup(function () {
        let filterKey = $('#phoneSearch').val().toUpperCase();
        let colNum = 8;
        doFilter(filterKey, colNum);
    });

    $('#emailSearch').keyup(function () {
        let filterKey = $('#emailSearch').val().toUpperCase();
        let colNum = 9;
        doFilter(filterKey, colNum);
    });
}

// calling filter func with resetting serial
function doFilter(filterKey, colNum) {
    filterRaw(filterKey, colNum);
    resetSerial();
    removeOtherFilter(colNum);
}

// removing other filter key
function removeOtherFilter(colNum) {
    $('#searchRow').find('th').each(function (index, td) {
        if (index + 1 != colNum) {
            $(td).find('input').val("");
        }
    });
}

// filter/search
function filterRaw(filterKey, colNum) {
    // Loop through all tr items, and hide those who don't match the search query
    $('tbody tr').each(function () {
        let tdValue = $(this).find("td:eq(" + colNum + ")").text();

        if (tdValue.toUpperCase().indexOf(filterKey) > -1) {
            $(this).css('display', '');
        } else {
            $(this).css('display', 'none');
        }
    });
}

// resetting serial after searching
function resetSerial() {
    let idx = 1;
    $('tbody tr').each(function () {
        let serial = justifySerial(idx);

        if ($(this).css('display') != 'none') {
            $(this).find("td:eq(1)").text(serial);
            idx++;
        }
    });
}

// error field clearing
function errFieldClearing() {
    $('#username').keyup(function () {
        $('#errUsername').text("");
    });
    $('#password').keyup(function () {
        $('#errPassword').text("");
    });
    $('#name').keyup(function () {
        $('#errName').text("");
    });
    $('#session').keyup(function () {
        $('#errSession').text("");
    });
    $('#reg').keyup(function () {
        $('#errReg').text("");
    });
    $('#roll').keyup(function () {
        $('#errRoll').text("");
    });
    $('#passingYear').keyup(function () {
        $('#errPassingYear').text("");
    });
    $('#email').keyup(function () {
        $('#errEmail').text("");
    });
}


// edit icon click function
function editItem(item) {
    $('#addModalLabel').text("Edit Information");
    $('#btnNo').text("Cancel");
    $('#addSubmit').val("Update");
    $('#addMessage').text("");
    clearAllErrorText();

    let sID = item.parentElement.parentElement.cells[0].innerText.trim();
    // let sAvatarURL = item.parentElement.parentElement.cells[2].innerText.trim();
    let sName = item.parentElement.parentElement.cells[3].innerText.trim();
    let sSession = item.parentElement.parentElement.cells[4].innerText.trim();
    let sReg = item.parentElement.parentElement.cells[5].innerText.trim();
    let sRoll = item.parentElement.parentElement.cells[6].innerText.trim();
    let sPassingYear = item.parentElement.parentElement.cells[7].innerText.trim();
    let sPhone = item.parentElement.parentElement.cells[8].innerText.trim();
    let sEmail = item.parentElement.parentElement.cells[9].innerText.trim();

    $('#sid').text(sID);
    $('#name').val(sName);
    $('#session').val(sSession);
    $('#reg').val(sReg);
    $('#roll').val(sRoll);
    $('#passingYear').val(sPassingYear);
    $('#phone').val(sPhone);
    $('#email').val(sEmail);
    // $('#avatar').val(sAvatarURL);

    // console.log(sID);
}

// delete icon click function
function deleteItem(item) {
    let sID = item.parentElement.parentElement.cells[0].innerText.trim();
    $('#sid').text(sID);
    $('#deleteMessage').text("");
}

function clearAllErrorText() {
    $('#errUsername').text("");
    $('#errPassword').text("");
    $('#errName').text("");
    $('#errSession').text("");
    $('#errReg').text("");
    $('#errRoll').text("");
    $('#errPassingYear').text("");
    $('#errEmail').text("");
}

function deleteRequest() {
    $('#deleteSubmit').prop('disabled', true);
    $('#deleteSubmit').val("Deleting...");

    let formData = $('#deleteModal form').serializeArray();
    // console.log(formData);

    // sending ajax post request
    let request = $.ajax({
        async: true,
        type: "DELETE",
        url: "/api/student/" + $('#sid').text().trim(),
    });
    request.done((response) => {
        // console.log(response)
        $('#deleteMessage').text(response.message);

        if (response.status == "error") {
            $('#deleteMessage').css('color', '#ff4d4d');
        } else {
            $('#deleteMessage').css('color', '#26C281');
            setTimeout(() => {
                location.reload();
            }, 500);
        }
    });
    request.fail((response) => {
        $('#deleteMessage').text(response.responseJSON.message);    // if response code in not 200(leave this for now)
        $('#deleteMessage').css('color', '#ff4d4d');
    });
    request.always(() => {
        $('#deleteSubmit').prop('disabled', false);
        $('#deleteSubmit').val("Delete");
    });
}

function addRequest() {
    let res = addFormValidate();
    if (!res) {
        return false;
    }

    let formArray = $('#addModal form').serializeArray();
    // console.log(formArray);

    let apiURL = "/api/student";    // add api url
    let apiMethod = "POST";    // add api url
    let todo = $('#addSubmit').val();
    let responseMessageText = "added";

    if (todo == "Add") {
        $('#addSubmit').prop('disabled', true);
        $('#addSubmit').val("Adding...");
    } else if (todo == "Update") {
        formArray.push({ name: 'id', value: $('#sid').text().trim() });
        apiURL = "/api/student/" + $('#sid').text().trim();
        apiMethod = "PUT";
        responseMessageText = "updated";

        $('#addSubmit').prop('disabled', true);
        $('#addSubmit').val("Updating...");
    }

    let formData = formArraytoJSON(formArray);  // converting form data to json obj
    // console.log(formData);

    // sending ajax post request
    let request = $.ajax({
        async: true,
        type: apiMethod,
        url: apiURL,
        data: formData,
        dataType: 'json',
    });
    request.done(function (response) {
        // console.log(response)

        if (response.status == "error") {
            if (response.type == "name") {
                $('#errName').text(response.message);
            } else if (response.type == "session") {
                $('#errSession').text(response.message);
            } else if (response.type == "reg") {
                $('#errReg').text(response.message);
            } else if (response.type == "roll") {
                $('#errRoll').text(response.message);
            } else if (response.type == "passingYear") {
                $('#errPassingYear').text(response.message);
            } else { // email error or any other errors
                $('#errEmail').text(response.message);
            }
        } else {
            $('#addMessage').text("data " + responseMessageText + " successfully");
            $('#addMessage').css('color', '#26C281');
            setTimeout(() => {
                location.reload();
            }, 500);
        }
    });
    request.fail(function (response) {
        $('#addMessage').text("something went wrong");
        $('#addMessage').css('color', '#ff4d4d');
        console.log(response)
    });
    request.always(function () {
        if (todo == "Add") {
            $('#addSubmit').prop('disabled', false);
            $('#addSubmit').val("Add");
        } else {
            $('#addSubmit').prop('disabled', false);
            $('#addSubmit').val("Update");
        }
    });
}

function addFormValidate() {
    // taking care of name
    if ($('#name').val().trim().length == 0) {
        $('#name').val("");
        $('#errname').text("name should no be empty");
        return false;   // cancel submission
    }
    // taking care of session
    if ($('#session').val().trim().length == 0) {
        $('#session').val("");
        $('#errSession').text("session should no be empty");
        return false;   // cancel submission
    }
    // taking care of reg
    if ($('#reg').val().trim().length == 0) {
        $('#reg').val("");
        $('#errReg').text("reg. no. should no be empty");
        return false;   // cancel submission
    }
    // taking care of roll
    if ($('#roll').val().trim().length == 0) {
        $('#roll').val("");
        $('#errRoll').text("class roll should no be empty");
        return false;   // cancel submission
    }
    // taking care of passingYear
    if ($('#passingYear').val().trim().length == 0) {
        $('#passingYear').val("");
        $('#errPassingYear').text("passing year should no be empty");
        return false;   // cancel submission
    }

    return true
}

function formArraytoJSON(formArray) {
    let JSONObj = {};
    for (var i = 0; i < formArray.length; i++) {
        JSONObj[formArray[i]['name']] = formArray[i]['value'];
    }
    return JSON.stringify(JSONObj);
}

function loginRequest() {
    if ($('#username').val().trim().length == 0) {
        $('#username').val("");
        $('#errUsername').text("username should no be empty");
        return false;   // cancel submission
    }

    $('#loginSubmit').prop('disabled', true);
    $('#loginSubmit').val("Logging in...");

    let formArray = $('#loginModal form').serializeArray();
    let formData = formArraytoJSON(formArray);
    // console.log(formData);

    // sending ajax post request
    let request = $.ajax({
        async: true,
        type: "POST",
        url: "/auth/login",
        data: formData,
        dataType: 'json',
    });
    request.done(function (response) {
        // console.log(response)

        if (response.status == "error") {
            if (response.message == "username not found") {
                $('#errUsername').text(response.message);
            } else {
                $('#errPassword').text(response.message);
            }
        } else {
            $('#loginMessage').text("login successful");
            $('#loginMessage').css('color', '#26C281');
            setTimeout(() => {
                location.reload();
            }, 500);
        }
    });
    request.fail(function (response) {
        console.log("something went wrong" + response)
    });
    request.always(function () {
        $('#loginSubmit').prop('disabled', false);
        $('#loginSubmit').val("Login");
    });
}

function getStudentList() {
    // sending ajax post request
    let request = $.ajax({
        async: true,
        type: "GET",
        url: "/api/student",
    });
    request.done(function (response) {
        displayList(response.data);
    });
    request.fail(function (response) {
        console.log("something went wrong" + response)
    });
    request.always(function () {
        // console.log("always")
    });
}

function displayList(list) {
    // removing current existing rows
    let rowSize = $("#table tbody tr").length;
    for (let i = 0; i < rowSize; i++) {
        $('.tableRow').remove();
    }

    $('#loadingGif').css("display", "none");
    if (list.length == null || list.length == 0) {
        $('#notFound').css('display', 'block'); //if no student found
    } else {
        $('#notFound').css('display', 'none');

        // adding new rows
        for (let i = 0; i < list.length; i++) {
            let sn = i + 1;
            let serial = justifySerial(sn);

            let avatarURL = "../assets/images/avatar.png"   //blank avatar link
            if (list[i].avatarURL.length > 0) {
                avatarURL = list[i].avatarURL;
            }
            let tableData = `<tr class="tableRow">
                        <td class="align-middle" id="sID" style="display: none;">`+ list[i].id + `</td>
                        <td class="align-middle">`+ serial + `</td>
                        <td class="align-middle p-1"><img src="`+ avatarURL + `" alt="" srcset="" class="avatar"></td>
                        <td class="align-middle">`+ list[i].name + `</td>
                        <td class="align-middle">`+ list[i].session + `</td>
                        <td class="align-middle">`+ list[i].reg + `</td>
                        <td class="align-middle">`+ list[i].roll + `</td>
                        <td class="align-middle">`+ list[i].passingYear + `</td>
                        <td class="align-middle">`+ list[i].phone + `</td>
                        <td class="align-middle">`+ list[i].email + `</td>`

            if ($('#loggedIn').text().trim() == "true") {
                tableData += `<td class="align-middle editIcon">
                                <img src="../assets/images/edit.png" alt="" srcset="" title="Edit" onclick="editItem(this)" data-bs-toggle="modal" data-bs-target="#addModal">
                                <img src="../assets/images/delete.png" alt="" srcset="" class="ms-3" title="Delete" onclick="deleteItem(this)" data-bs-toggle="modal" data-bs-target="#deleteModal">
                            </td>`
            }
            tableData += `</tr>`

            $('#table').append(tableData);
        }
    }
}

function justifySerial(sn) {
    let serial = "";
    if (sn < 10) {
        serial = "00" + sn.toString();
    } else if (sn < 100) {
        serial = "0" + sn.toString();
    } else {
        serial = sn;
    }

    return serial;
}