import static org.junit.Assert.assertThat;
import static org.hamcrest.Matchers.is;

import java.util.Objects;
import java.util.function.BiFunction;
import java.util.function.Predicate;

import org.junit.Test;

/**
 * https://gist.github.com/tamada/930bff6d3a72da2533eacf3a833b7563
 */
public class PairTest {
    private BiFunction<Integer, String, Integer> leftGetter = (l, r) -> l;
    private BiFunction<Integer, String, String> rightGetter = (l, r) -> r;

    @Test
    public void testUnify() {
        var pair = Pair.of(10, "ten");
        assertThat(pair.unify(leftGetter), is(10));
        assertThat(pair.unify(rightGetter), is("ten"));
    }

    @Test
    public void testSwap() {
        var pair = Pair.of("ten", 10).swap();
        assertThat(pair.unify(leftGetter), is(10));
        assertThat(pair.unify(rightGetter), is("ten"));
    }

    @Test
    public void testMap() {
        var pair = Pair.of(10, "ten")
                .map(l -> l * 10, r -> r + r);
        assertThat(pair.unify(leftGetter), is(100));
        assertThat(pair.unify(rightGetter), is("tenten"));
    }

    @Test
    public void testTest() {
        var pair = Pair.of(10, "ten");
        assertThat(pair.equals(Pair.of(5 * 2, new String("ten"))), is(true));
    }

    @Test
    public void testTest2() {
        var pair = Pair.of(10, "ten");
        assertThat(pair.test(predicate(10), predicate("ten")), is(Pair.of(true, true)));
        assertThat(pair.test(predicate(10), predicate("one")), is(Pair.of(true, false)));
        assertThat(pair.test(predicate(90), predicate("ten")), is(Pair.of(false, true)));
        assertThat(pair.test(predicate(90), predicate("one")), is(Pair.of(false, false)));
    }

    private <T> Predicate<T> predicate(T value) {
        return v -> Objects.equals(value, v);
    }

    @Test
    public void testEquals() {
        var pair = Pair.of(10, "ten");
        assertThat(pair.equals(new Object()), is(false));
        assertThat(pair.equals(Pair.of(5, "ten")), is(false));
        assertThat(pair.equals(Pair.of(10, "notTen")), is(false));
        assertThat(pair.equals(Pair.of(5, "ten")), is(false));
        assertThat(pair.equals(Pair.of(10, "ten")), is(true));
    }

    @Test
    public void testToString() {
        var pair = Pair.of(10, "ten");
        assertThat(pair.toString(), is("(10, ten)"));
    }

    @Test(expected=NullPointerException.class)
    public void testInstantiateFailedLeftWasNull() {
        Pair.of(null, "ten");
    }

    @Test(expected=NullPointerException.class)
    public void testInstantiateFailedRightWasNull() {
        Pair.of(10, null);
    }
}
