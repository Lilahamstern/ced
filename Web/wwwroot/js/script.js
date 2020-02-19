$(document).ready(function () {
    $('#projects').DataTable({
        "paging": false,
        "searching": false,
        "aaSorting": [],
        columnDefs: [{
            orderable: false,
            targets: [0, 7]
        }],
        "info": false,
    });
    $('.dataTables_length').addClass('bs-select');
});