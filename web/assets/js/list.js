$(document).ready(function() {
    APIURL = $("body").data("api-host")
    cochonou_domain = $("body").data("cochonou-domain")

    copy = new Clipboard('.copy');
    copy.on('success', function(e) {
        sweetAlert({
              title: "Copied !",
              timer: 1000,
              type: 'success',
              showConfirmButton: false
          });
    });
    $.getJSON( APIURL + "/redirections", function( data ) {
        $.each(data, function(key, redir) {
            $("#link-list").append(`
                <li class="link" data-link="http://`+redir.sub_domain+`.`+cochonou_domain+`">
                    `+redir.sub_domain+`
                    <div class="actions">
                        <button class="preview">
                        <svg height="15" width="15">
                            <use xlink:href="#preview"></use>
                        </svg>
                        </button>
                        <button class="copy" data-clipboard-action="copy" data-clipboard-text="http://`+redir.sub_domain+`.`+cochonou_domain+`">
                        <svg height="15" width="15">
                            <use xlink:href="#copy"></use>
                        </svg>
                        </button>
                        <button class="delete">
                        <svg height="15" width="15">
                            <use xlink:href="#delete"></use>
                        </svg>
                        </button>
                    </div>
                    <hr>
                    <img class="preview" src = "" />
                </li>
            `);
        });
    });

    $('#link-list').on('click', '.link .preview', function(e){
        linkLi = $(this).parent().parent()
        linkLi.children('.preview').attr('src', linkLi.data('link'))
        linkLi.toggleClass('active')
    })
    $('#link-list').on('click', '.link', function(e){
        if(e.target != this) return;
        linkLi = $(this)
        linkLi.children('.preview').attr('src', linkLi.data('link'))
        linkLi.toggleClass('active')
    })
});
