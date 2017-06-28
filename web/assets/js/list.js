$( document ).ready(function() {
    APIURL = $("body").data("api-host")
    cochonou_domain = $("body").data("cochonou-domain")
    $.getJSON( APIURL + "/redirections", function( data ) {
        $.each(data, function(key, redir) {
            $("#link-list").append(`
                <li>
                    <a href="http://`+redir.sub_domain+`.`+cochonou_domain+`">
                        http://`+redir.sub_domain+`.`+cochonou_domain+`
                    </a>
                    <div class="actions">
                        <button class="preview">
                        <svg height="15" width="15">
                            <use xlink:href="#preview"></use>
                        </svg>
                        </button>
                        <button class="copy">
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
                    Preview link
                </li>
            `);
        });
    });
});