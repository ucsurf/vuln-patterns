fn a() {
    // ruleid: ammonia-svg-animation-tags
    let c = ammonia::Builder::default().add_tags(&["svg", "animateTransform"]).clean(x);
}
fn b(b: &mut ammonia::Builder) {
    // ruleid: ammonia-svg-animation-tags
    b.add_tags(&["set"]);
}
fn ok() {
    // ok: ammonia-svg-animation-tags
    let c = ammonia::Builder::default().add_tags(&["svg", "circle"]).clean(x);
}
