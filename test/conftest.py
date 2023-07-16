import os
import subprocess
import tempfile

import py
import pytest


# ######## #
# Fixtures #
# ######## #
def get_tmpdir(request):
    tmpdir = py.path.local(tempfile.mkdtemp())
    request.addfinalizer(lambda: tmpdir.remove(rec=1))
    return str(tmpdir)


@pytest.fixture(scope="session")
def session_tmpdir(request):
    return get_tmpdir(request)


@pytest.fixture(scope="session")
def test_app(session_tmpdir):
    app_path = os.path.relpath(os.path.join(os.path.dirname(__file__), ".."))
    build_path = os.path.join(session_tmpdir, "test_app")

    build_env = os.environ.copy()
    build_env["TEST_APP_BUILD_PATH"] = build_path

    process = subprocess.run(["mage", "build"], cwd=app_path,
                             env=build_env)
    assert process.returncode == 0, "Failed to build app"

    return build_path
