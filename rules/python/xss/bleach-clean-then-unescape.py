import html, bleach, nh3
def a(x):
    # ruleid: sanitizer-output-mutated-after-clean
    return html.unescape(bleach.clean(x))
def b(x):
    c = nh3.clean(x)
    # ruleid: sanitizer-output-mutated-after-clean
    return html.unescape(c)
def ok(x):
    # ok: sanitizer-output-mutated-after-clean
    return bleach.clean(html.unescape(x))
