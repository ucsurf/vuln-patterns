fn a(p: &Path, m: Permissions) -> io::Result<()> {
    let f = File::create(p)?;
    // ruleid: create-then-set-permissions
    fs::set_permissions(p, m)?;
    Ok(())
}
fn b(p: &Path, m: Permissions) -> io::Result<()> {
    let f = File::create(p)?;
    // ruleid: create-then-set-permissions
    f.set_permissions(m)?;
    Ok(())
}
fn ok(p: &Path) -> io::Result<()> {
    // ok: create-then-set-permissions
    let f = OpenOptions::new().write(true).create_new(true).mode(0o600).open(p)?;
    Ok(())
}
