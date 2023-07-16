import subprocess

import pytest


@pytest.mark.parametrize("args, expected_rc, expected_out", [
    pytest.param(
        ["add", "1", "2"],
        0,
        "3\n",
    ),
    pytest.param(
        ["mul", "14", "10"],
        0,
        "140\n"
    ),
    pytest.param(
        ["add", "150"],
        0,
        "150\n"
    ),
    pytest.param(
        ["mul", "6996"],
        0,
        "0\n",
    ),
    pytest.param(
        ["kek", "1", "2"],
        2,
        "",
    ),
])
def test_app(test_app, args, expected_rc, expected_out):
    cmd = [test_app] + args

    process = subprocess.Popen(
        cmd,
        stderr=subprocess.STDOUT,
        stdout=subprocess.PIPE,
        text=True
    )
    rc = process.wait()

    assert rc == expected_rc

    output = process.stdout.read()
    if rc == 0:
        assert output == expected_out
