import html, bleach

def render(comment):
    return html.unescape(bleach.clean(comment))
