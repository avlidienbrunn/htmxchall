htmx.defineExtension('hx-purify', {
    transformResponse : function(text, xhr, elt) {return DOMPurify.sanitize(text);},
});